package httpapi

import (
	"database/sql"
	"io"

	"github.com/korylprince/httputil/auth"
	"github.com/korylprince/httputil/session"
	"github.com/korylprince/printer-manager/sync"
)

//Server represents shared resources
type Server struct {
	db           *sql.DB
	manager      *sync.Manager
	auth         auth.Auth
	sessionStore session.Store
	output       io.Writer
}

//NewServer returns a new server with the given resources
func NewServer(db *sql.DB, manager *sync.Manager, auth auth.Auth, sessionStore session.Store, output io.Writer) *Server {
	return &Server{db: db, manager: manager, auth: auth, sessionStore: sessionStore, output: output}
}
