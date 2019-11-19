package sync

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	ldap "github.com/go-ldap/ldap/v3"
	"github.com/gofrs/uuid"
	auth "github.com/korylprince/go-ad-auth/v3"
	"github.com/korylprince/printer-manager/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type user struct {
	ID          string
	LocalID     string
	DisplayName string
	Username    string
	Groups      []*group
}

type group struct {
	ID          string
	LocalID     string
	DisplayName string
	Users       []*user
}

func bind(config *auth.Config, username, password string) (*auth.Conn, error) {
	upn, err := config.UPN(username)
	if err != nil {
		return nil, fmt.Errorf("Unable to generate userPrincipalName: %v", err)
	}

	conn, err := config.Connect()
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to Active Directory: %v", err)
	}

	ok, err := conn.Bind(upn, password)
	if err != nil {
		return nil, fmt.Errorf("Error attempting to bind to Active Directory: %v", err)
	}
	if !ok {
		return nil, errors.New("Error attempting to bind to Active Directory: Invalid credentials")
	}

	return conn, nil
}

func pagedSearch(conn *auth.Conn, filter string, attrs []string) ([]*ldap.Entry, error) {
	search := ldap.NewSearchRequest(
		conn.Config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		filter,
		attrs,
		nil,
	)
	result, err := conn.Conn.SearchWithPaging(search, 1000)
	if err != nil {
		return nil, fmt.Errorf(`Search error "%s": %v`, filter, err)
	}

	return result.Entries, nil
}

func getADUsersAndGroups(conn *auth.Conn, userFilter, userAttr, groupFilter string) ([]*user, []*group, error) {
	userEntries, err := pagedSearch(conn, userFilter, []string{"displayName", "sAMAccountName", "memberOf", userAttr})
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to search for users: %v", err)
	}

	groupEntries, err := pagedSearch(conn, groupFilter, []string{"cn", "objectGUID"})
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to search for groups: %v", err)
	}

	var groups []*group
	var groupMap = make(map[string]*group)
	for _, e := range groupEntries {
		id, err := uuid.FromBytes(e.GetRawAttributeValue("objectGUID"))
		if err != nil {
			return nil, nil, fmt.Errorf("Unable to parse UUID: %v", err)
		}
		g := &group{
			DisplayName: e.GetAttributeValue("cn"),
			LocalID:     id.String(),
			Users:       make([]*user, 0),
		}
		groups = append(groups, g)
		groupMap[e.DN] = g
	}

	var users []*user
	for _, e := range userEntries {
		var groups []*group
		for _, dn := range e.GetAttributeValues("memberOf") {
			if g, ok := groupMap[dn]; ok {
				groups = append(groups, g)
			}
		}
		u := &user{
			DisplayName: e.GetAttributeValue("displayName"),
			LocalID:     e.GetAttributeValue(userAttr),
			Username:    e.GetAttributeValue("sAMAccountName"),
			Groups:      groups,
		}
		users = append(users, u)
		for _, g := range groups {
			g.Users = append(g.Users, u)
		}
	}

	var filteredGroups []*group
	for _, g := range groups {
		if len(g.Users) > 0 {
			filteredGroups = append(filteredGroups, g)
		}
	}

	return users, filteredGroups, nil
}

func getDBUsersAndGroups(pdb *sql.Tx) ([]*user, []*group, error) {
	dbUsers, err := db.Users(qm.Load("Groups")).All(context.Background(), pdb)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to query users: %v", err)
	}

	dbGroups, err := db.Groups().All(context.Background(), pdb)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to query groups: %v", err)
	}

	var groups []*group
	groupMap := make(map[string]*group)
	for _, dg := range dbGroups {
		g := &group{
			ID:          dg.ID,
			LocalID:     dg.LocalID,
			DisplayName: dg.DisplayName,
			Users:       make([]*user, 0),
		}
		groups = append(groups, g)
		groupMap[g.LocalID] = g
	}

	var users []*user
	for _, du := range dbUsers {
		var groups []*group
		for _, dg := range du.R.Groups {
			groups = append(groups, groupMap[dg.LocalID])
		}

		u := &user{
			ID:          du.ID,
			LocalID:     du.LocalID,
			DisplayName: du.DisplayName,
			Username:    du.Username,
			Groups:      groups,
		}

		for _, g := range groups {
			g.Users = append(g.Users, u)
		}

		users = append(users, u)
	}

	return users, groups, nil
}

func syncADToDB(tx *sql.Tx, adUsers, dbUsers []*user, adGroups, dbGroups []*group) error {
	dbUserMap := make(map[string]*user)
	for _, du := range dbUsers {
		dbUserMap[du.LocalID] = du
	}
	adUserMap := make(map[string]*user)
	for _, au := range adUsers {
		adUserMap[au.LocalID] = au
	}

	var createUsers []*user
	var updateUsers []*user
	var deleteUsers []*user

	for _, au := range adUsers {
		du, ok := dbUserMap[au.LocalID]
		if !ok {
			createUsers = append(createUsers, au)
			continue
		}
		if au.DisplayName != du.DisplayName || au.Username != du.Username {
			updateUsers = append(updateUsers, au)
		}
		adUserMap[du.LocalID].ID = du.ID
	}

	for _, du := range dbUsers {
		if _, ok := adUserMap[du.LocalID]; !ok {
			deleteUsers = append(deleteUsers, du)
		}
	}

	dbGroupMap := make(map[string]*group)
	for _, dg := range dbGroups {
		dbGroupMap[dg.LocalID] = dg
	}
	adGroupMap := make(map[string]*group)
	for _, ag := range adGroups {
		adGroupMap[ag.LocalID] = ag
	}

	var createGroups []*group
	var updateGroups []*group
	var deleteGroups []*group

	for _, ag := range adGroups {
		dg, ok := dbGroupMap[ag.LocalID]
		if !ok {
			createGroups = append(createGroups, ag)
			continue
		}
		if ag.DisplayName != dg.DisplayName {
			updateGroups = append(updateGroups, ag)
		}
		adGroupMap[dg.LocalID].ID = dg.ID
	}

	for _, dg := range dbGroups {
		if _, ok := adGroupMap[dg.LocalID]; !ok {
			deleteGroups = append(deleteGroups, dg)
		}
	}

	for _, u := range createUsers {
		du := &db.User{
			LocalID:     u.LocalID,
			DisplayName: u.DisplayName,
			Username:    u.Username,
		}
		log.Printf("Creating User: %s\n", u.DisplayName)
		if err := du.Insert(context.Background(), tx, boil.Blacklist(db.UserColumns.ID)); err != nil {
			return fmt.Errorf("Unable to create user: %v", err)
		}
		adUserMap[du.LocalID].ID = du.ID
	}

	for _, u := range updateUsers {
		du := &db.User{
			ID:          u.ID,
			LocalID:     u.LocalID,
			DisplayName: u.DisplayName,
			Username:    u.Username,
		}
		log.Printf("Updating User: %s\n", u.DisplayName)
		if _, err := du.Update(context.Background(), tx, boil.Blacklist(db.UserColumns.ID)); err != nil {
			return fmt.Errorf("Unable to update user: %v", err)
		}
	}

	for _, u := range deleteUsers {
		du := &db.User{
			ID: u.ID,
		}
		log.Printf("Deleting User: %s\n", u.DisplayName)
		if _, err := du.Delete(context.Background(), tx); err != nil {
			return fmt.Errorf("Unable to delete user: %v", err)
		}
	}

	for _, g := range createGroups {
		dg := &db.Group{
			LocalID:     g.LocalID,
			DisplayName: g.DisplayName,
		}
		log.Printf("Creating Group: %s\n", g.DisplayName)
		if err := dg.Insert(context.Background(), tx, boil.Blacklist(db.GroupColumns.ID)); err != nil {
			return fmt.Errorf("Unable to create group: %v", err)
		}
		adGroupMap[dg.LocalID].ID = dg.ID
	}

	for _, g := range updateGroups {
		dg := &db.Group{
			ID:          g.ID,
			LocalID:     g.LocalID,
			DisplayName: g.DisplayName,
		}
		log.Printf("Updating Group: %s\n", g.DisplayName)
		if _, err := dg.Update(context.Background(), tx, boil.Blacklist(db.GroupColumns.ID)); err != nil {
			return fmt.Errorf("Unable to update group: %v", err)
		}
	}

	for _, g := range deleteGroups {
		dg := &db.Group{
			ID: g.ID,
		}
		log.Printf("Deleting Group: %s\n", g.DisplayName)
		if _, err := dg.Delete(context.Background(), tx); err != nil {
			return fmt.Errorf("Unable to delete group: %v", err)
		}
	}

	adAssignments := make(map[[2]string]struct{})
	for _, ag := range adGroups {
		for _, au := range ag.Users {
			adAssignments[[2]string{ag.LocalID, au.LocalID}] = struct{}{}
		}
	}

	dbAssignments := make(map[[2]string]struct{})
	for _, dg := range dbGroups {
		for _, du := range dg.Users {
			dbAssignments[[2]string{dg.LocalID, du.LocalID}] = struct{}{}
		}
	}

	for aid := range adAssignments {
		if _, ok := dbAssignments[aid]; !ok {
			dg := &db.Group{ID: adGroupMap[aid[0]].ID}
			du := &db.User{ID: adUserMap[aid[1]].ID}
			log.Printf("Assigning Group: %s, User: %s\n", adGroupMap[aid[0]].DisplayName, adUserMap[aid[1]].DisplayName)
			if err := dg.AddUsers(context.Background(), tx, false, du); err != nil {
				return fmt.Errorf("Unable to assign group user: %v", err)
			}
		}
	}

	for did := range dbAssignments {
		if _, ok := adAssignments[did]; !ok {
			dg := &db.Group{ID: dbGroupMap[did[0]].ID}
			du := &db.User{ID: dbUserMap[did[1]].ID}
			log.Printf("Unassigning Group: %s, User: %s\n", dbGroupMap[did[0]].DisplayName, dbUserMap[did[1]].DisplayName)
			if err := dg.RemoveUsers(context.Background(), tx, du); err != nil {
				return fmt.Errorf("Unable to unassign group user: %v", err)
			}
		}
	}

	return nil
}

func sync(config *Config, tx *sql.Tx) error {
	conn, err := bind(&auth.Config{
		Server:   config.Server,
		Port:     config.Port,
		Security: config.Security,
		BaseDN:   config.BaseDN,
	}, config.BindUsername, config.BindPassword)
	if err != nil {
		return fmt.Errorf("Unable to bind to Active Directory: %v", err)
	}
	defer conn.Conn.Close()

	adUsers, adGroups, err := getADUsersAndGroups(conn, config.UserFilter, config.UserAttribute, config.GroupFilter)
	if err != nil {
		return fmt.Errorf("Unable to get AD users and groups: %v", err)
	}

	dbUsers, dbGroups, err := getDBUsersAndGroups(tx)
	if err != nil {
		return fmt.Errorf("Unable to get DB users and groups: %v", err)
	}

	if err = syncADToDB(tx, adUsers, dbUsers, adGroups, dbGroups); err != nil {
		return fmt.Errorf("Unable to complete sync: %v", err)
	}

	return nil
}
