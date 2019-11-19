// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package crud

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/gorilla/mux"
	"github.com/korylprince/httputil/jsonapi"
	"github.com/korylprince/printer-manager/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func CreateUser(r *http.Request, tx *sql.Tx) (int, interface{}) {
	user := new(db.User)
	if err := jsonapi.ParseJSONBody(r, user); err != nil {
		return http.StatusBadRequest, err
	}

	if err := user.Insert(r.Context(), tx, boil.Blacklist(
		db.UserColumns.ID,
	)); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates") {
			return http.StatusConflict, fmt.Errorf("Unable to insert duplicate User: %v", err)
		} else if strings.Contains(err.Error(), "value too long for type character varying") {
			return http.StatusBadRequest, fmt.Errorf("Unable to insert invalid User: %v", err)
		} else if strings.Contains(err.Error(), "violates check constraint") {
			return http.StatusBadRequest, fmt.Errorf("Unable to insert invalid User: %v", err)
		} else if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return http.StatusBadRequest, fmt.Errorf("Unable to insert invalid User: %v", err)
		} else if strings.Contains(err.Error(), "violates foreign key constraint") {
			return http.StatusBadRequest, fmt.Errorf("Unable to insert invalid User: %v", err)
		}
		return http.StatusInternalServerError, fmt.Errorf("Unable to insert User: %v", err)
	}

	return http.StatusOK, user
}

func ReadUser(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)
	id := vars[db.UserColumns.ID]

	user, err := db.FindUser(r.Context(), tx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return http.StatusNotFound, fmt.Errorf("User %s does not exist", id)
		}

		return http.StatusInternalServerError, fmt.Errorf("Unable to find User %s: %v", id, err)
	}

	return http.StatusOK, user
}

func UpdateUser(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)
	id := vars[db.UserColumns.ID]

	newUser := new(db.User)
	if err := jsonapi.ParseJSONBody(r, newUser); err != nil {
		return http.StatusBadRequest, err
	}

	user, err := db.FindUser(r.Context(), tx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return http.StatusNotFound, fmt.Errorf("User %s does not exist", id)
		}

		return http.StatusInternalServerError, fmt.Errorf("Unable to find User %s: %v", id, err)
	}

	user.LocalID = newUser.LocalID

	user.DisplayName = newUser.DisplayName

	user.Username = newUser.Username

	if _, err = user.Update(r.Context(), tx, boil.Blacklist(
		db.UserColumns.ID,
	)); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates") {
			return http.StatusConflict, fmt.Errorf("Unable to insert duplicate User: %v", err)
		} else if strings.Contains(err.Error(), "value too long for type character varying") {
			return http.StatusBadRequest, fmt.Errorf("Unable to insert invalid User: %v", err)
		} else if strings.Contains(err.Error(), "violates check constraint") {
			return http.StatusBadRequest, fmt.Errorf("Unable to insert invalid User: %v", err)
		} else if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return http.StatusBadRequest, fmt.Errorf("Unable to insert invalid User: %v", err)
		} else if strings.Contains(err.Error(), "violates foreign key constraint") {
			return http.StatusBadRequest, fmt.Errorf("Unable to insert invalid User: %v", err)
		}
		return http.StatusInternalServerError, fmt.Errorf("Unable to insert User: %v", err)
	}

	return http.StatusOK, user
}

func DeleteUser(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)
	id := vars[db.UserColumns.ID]

	user, err := db.FindUser(r.Context(), tx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return http.StatusNotFound, fmt.Errorf("User %s does not exist", id)
		}

		return http.StatusInternalServerError, fmt.Errorf("Unable to find User %s: %v", id, err)
	}

	groups, err := user.Groups().All(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to find User %s Groups: %v", id, err)
	}
	if len(groups) > 0 {
		return http.StatusConflict, fmt.Errorf("User %s still in use", id)
	}

	locations, err := user.Locations().All(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to find User %s Locations: %v", id, err)
	}
	if len(locations) > 0 {
		return http.StatusConflict, fmt.Errorf("User %s still in use", id)
	}

	if _, err = user.Delete(r.Context(), tx); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to delete User %s: %v", id, err)
	}

	return http.StatusOK, nil
}

func QueryUsers(r *http.Request, tx *sql.Tx) (int, interface{}) {
	params := r.URL.Query()
	var mods []qm.QueryMod

	if _, ok := params[db.UserColumns.LocalID]; ok {
		mods = append(mods, qm.Where("local_id LIKE ?", params.Get(db.UserColumns.LocalID)))
	}
	if _, ok := params[db.UserColumns.DisplayName]; ok {
		mods = append(mods, qm.Where("display_name LIKE ?", params.Get(db.UserColumns.DisplayName)))
	}
	if _, ok := params[db.UserColumns.Username]; ok {
		mods = append(mods, qm.Where("username LIKE ?", params.Get(db.UserColumns.Username)))
	}

	users, err := db.Users(mods...).All(r.Context(), tx)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return http.StatusNotFound, errors.New("No Users found matching query")
		}

		return http.StatusInternalServerError, fmt.Errorf("Unable to query Users: %v", err)
	}

	if users == nil {
		return http.StatusOK, db.UserSlice{}
	}

	return http.StatusOK, users
}

func RelateUserGroup(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)

	vars[db.UserColumns.ID] = vars["user_id"]
	delete(vars, "user_id")
	code, v := ReadUser(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	user := v.(*db.User)

	vars[db.UserColumns.ID] = vars["group_id"]
	delete(vars, "group_id")
	code, v = ReadGroup(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	group := v.(*db.Group)

	ok, err := user.Groups(qm.Where("id = ?", group.ID)).Exists(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to check if User related to Group: %v", err)
	}
	if ok {
		return http.StatusOK, nil
	}

	if err := user.AddGroups(r.Context(), tx, false, group); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to relate User to Group: %v", err)
	}

	return http.StatusOK, nil
}

func UnrelateUserGroup(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)

	vars[db.UserColumns.ID] = vars["user_id"]
	delete(vars, "user_id")
	code, v := ReadUser(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	user := v.(*db.User)

	vars[db.UserColumns.ID] = vars["group_id"]
	delete(vars, "group_id")
	code, v = ReadGroup(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	group := v.(*db.Group)

	ok, err := user.Groups(qm.Where("id = ?", group.ID)).Exists(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to check if User related to Group: %v", err)
	}
	if !ok {
		return http.StatusOK, nil
	}

	if err := user.RemoveGroups(r.Context(), tx, group); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to unrelate User to Group: %v", err)
	}

	return http.StatusOK, nil
}

func ReadUserGroups(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)

	vars[db.UserColumns.ID] = vars["user_id"]
	delete(vars, "user_id")
	code, v := ReadUser(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	user := v.(*db.User)

	groups, err := user.Groups().All(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to read User Groups: %v", err)
	}

	if groups == nil {
		return http.StatusOK, db.GroupSlice{}
	}

	return http.StatusOK, groups
}

func RelateUserLocation(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)

	vars[db.UserColumns.ID] = vars["user_id"]
	delete(vars, "user_id")
	code, v := ReadUser(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	user := v.(*db.User)

	vars[db.UserColumns.ID] = vars["location_id"]
	delete(vars, "location_id")
	code, v = ReadLocation(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	location := v.(*db.Location)

	ok, err := user.Locations(qm.Where("id = ?", location.ID)).Exists(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to check if User related to Location: %v", err)
	}
	if ok {
		return http.StatusOK, nil
	}

	if err := user.AddLocations(r.Context(), tx, false, location); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to relate User to Location: %v", err)
	}

	return http.StatusOK, nil
}

func UnrelateUserLocation(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)

	vars[db.UserColumns.ID] = vars["user_id"]
	delete(vars, "user_id")
	code, v := ReadUser(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	user := v.(*db.User)

	vars[db.UserColumns.ID] = vars["location_id"]
	delete(vars, "location_id")
	code, v = ReadLocation(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	location := v.(*db.Location)

	ok, err := user.Locations(qm.Where("id = ?", location.ID)).Exists(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to check if User related to Location: %v", err)
	}
	if !ok {
		return http.StatusOK, nil
	}

	if err := user.RemoveLocations(r.Context(), tx, location); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to unrelate User to Location: %v", err)
	}

	return http.StatusOK, nil
}

func ReadUserLocations(r *http.Request, tx *sql.Tx) (int, interface{}) {
	vars := mux.Vars(r)

	vars[db.UserColumns.ID] = vars["user_id"]
	delete(vars, "user_id")
	code, v := ReadUser(r, tx)
	if err, ok := v.(error); ok {
		return code, err
	}
	user := v.(*db.User)

	locations, err := user.Locations().All(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to read User Locations: %v", err)
	}

	if locations == nil {
		return http.StatusOK, db.LocationSlice{}
	}

	return http.StatusOK, locations
}
