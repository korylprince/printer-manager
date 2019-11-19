package sync

import (
	"errors"

	auth "github.com/korylprince/go-ad-auth/v3"
)

//Config represents the command line or environment variable configuration
type Config struct {
	Server        string
	Port          int
	Security      auth.SecurityType
	BaseDN        string
	BindUsername  string
	BindPassword  string
	UserFilter    string
	UserAttribute string
	GroupFilter   string
}

//Validate returns an error if the Config is not valid
func (c *Config) Validate() error {
	switch {
	case c.Server == "":
		return errors.New("Server must be set")
	case c.Port == 0:
		return errors.New("Port must be set")
	case c.BaseDN == "":
		return errors.New("BaseDN must be set")
	case c.BindUsername == "":
		return errors.New("BindUsername must be set")
	case c.BindPassword == "":
		return errors.New("BindPassword must be set")
	case c.UserFilter == "":
		return errors.New("UserFilter must be set")
	case c.UserAttribute == "":
		return errors.New("UserAttribute must be set")
	case c.GroupFilter == "":
		return errors.New("GroupFilter must be set")
	}
	return nil
}
