package main

import (
	"log"
	"strings"

	"github.com/kelseyhightower/envconfig"
	auth "github.com/korylprince/go-ad-auth/v3"
)

//Config represents options given in the environment
type Config struct {
	SessionExpiration int `default:"60"` //in minutes

	LDAPServer   string `required:"true"`
	LDAPPort     int    `required:"true" default:"389"`
	LDAPBaseDN   string `required:"true"`
	LDAPGroup    string `required:"true"`
	LDAPSecurity string `required:"true" default:"none"`

	LDAPBindUsername string `required:"true"`
	LDAPBindPassword string `required:"true"`
	UserFilter       string `required:"true" default:"(&(objectClass=user)(!userAccountControl:1.2.840.113556.1.4.803:=2))"`
	UserAttribute    string `required:"true" default:"employeeID"`
	GroupFilter      string `required:"true" default:"(objectClass=group)"`
	SyncInterval     int    `default:"30"` //in minutes

	SQLDSN string `required:"true"`

	ListenAddr string `required:"true" default:":80"` //addr format used for net.Dial; required
	Prefix     string //url prefix to mount api to without trailing slash
}

//SecurityType returns the auth.SecurityType for the config
func (c *Config) SecurityType() auth.SecurityType {
	switch strings.ToLower(c.LDAPSecurity) {
	case "", "none":
		return auth.SecurityNone
	case "tls":
		return auth.SecurityTLS
	case "starttls":
		return auth.SecurityStartTLS
	default:
		log.Fatalln("Invalid PRINTERMANAGER_LDAPSECURITY:", config.LDAPSecurity)
	}
	panic("unreachable")
}

var config = &Config{}

func init() {
	err := envconfig.Process("PRINTERMANAGER", config)
	if err != nil {
		log.Fatalln("Error reading configuration from environment:", err)
	}
}
