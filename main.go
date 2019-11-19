package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	adauth "github.com/korylprince/go-ad-auth/v3"
	"github.com/korylprince/httputil/auth/ad"
	"github.com/korylprince/httputil/session/memory"
	"github.com/korylprince/printer-manager/httpapi"
	"github.com/korylprince/printer-manager/sync"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", config.SQLDSN)
	if err != nil {
		log.Fatalln("Unable to open database:", err)
	}

	c := &sync.Config{
		Server:        config.LDAPServer,
		Port:          config.LDAPPort,
		Security:      config.SecurityType(),
		BaseDN:        config.LDAPBaseDN,
		BindUsername:  config.LDAPBindUsername,
		BindPassword:  config.LDAPBindPassword,
		UserFilter:    config.UserFilter,
		UserAttribute: config.UserAttribute,
		GroupFilter:   config.GroupFilter,
	}

	manager, err := sync.NewManager(c, db, time.Minute*time.Duration(config.SyncInterval))
	if err != nil {
		log.Fatalln("Unable to start manager:", err)
	}
	manager.Start()

	adConfig := &adauth.Config{
		Server:   config.LDAPServer,
		Port:     config.LDAPPort,
		BaseDN:   config.LDAPBaseDN,
		Security: config.SecurityType(),
	}
	auth := ad.New(adConfig, nil, []string{config.LDAPGroup})

	s := httpapi.NewServer(db, manager, auth, memory.New(time.Minute*time.Duration(config.SessionExpiration)), os.Stdout)

	log.Println("Listening on:", config.ListenAddr)
	log.Println(http.ListenAndServe(config.ListenAddr, http.StripPrefix(config.Prefix, s.Router())))
}
