package httpapi

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/korylprince/printer-manager/db"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

//Printer represents a network printer
type Printer struct {
	ID       string      `json:"id"`
	Hostname string      `json:"hostname"`
	Name     string      `json:"name"`
	Location string      `json:"location"`
	Driver   interface{} `json:"driver"`
}

func newPrinter(p *db.Printer) *Printer {
	var extra []byte
	if p.DriverExtra.Valid {
		extra = p.DriverExtra.JSON
	}

	driver, err := mergeJSON(p.R.Model.Driver, extra)
	if err != nil {
		log.Println("WARNING: Unable to merge JSON:", err)
		driver = p.R.Model.Driver
	}

	newP := &Printer{
		ID:       p.ID,
		Hostname: p.Hostname,
		Name:     fmt.Sprintf("%s %s %s", p.R.Location.Name, p.R.Model.R.Manufacturer.Name, p.R.Model.Name),
		Location: fmt.Sprintf("%s - %s", p.R.Location.R.Building.Name, p.R.Location.Name),
		Driver:   driver,
	}
	return newP
}

func readUserPrinters(r *http.Request, tx *sql.Tx) (int, interface{}) {
	username, ok := mux.Vars(r)["username"]
	if !ok || username == "" {
		return http.StatusBadRequest, errors.New("Request missing username")
	}
	user, err := db.Users(qm.Where("username = ?", username)).One(r.Context(), tx)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return http.StatusNotFound, fmt.Errorf("User %s does not exist", username)
		}

		return http.StatusInternalServerError, fmt.Errorf("Unable to find User %s: %v", username, err)
	}

	printers := make([]*Printer, 0)
	seen := make(map[string]struct{})

	locations, err := user.Locations(qm.Load("Printers.Model.Manufacturer"), qm.Load("Building")).All(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to query User Locations: %v", err)
	}
	for _, l := range locations {
		for _, p := range l.R.Printers {
			if _, ok := seen[p.ID]; !ok {
				printers = append(printers, newPrinter(p))
				seen[p.ID] = struct{}{}
			}
		}
	}

	groups, err := user.Groups(qm.Load("Locations.Printers.Model.Manufacturer"), qm.Load("Locations.Building")).All(r.Context(), tx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to query User Groups: %v", err)
	}
	for _, g := range groups {
		for _, l := range g.R.Locations {
			for _, p := range l.R.Printers {
				if _, ok := seen[p.ID]; !ok {
					printers = append(printers, newPrinter(p))
					seen[p.ID] = struct{}{}
				}
			}
		}
	}

	return http.StatusOK, printers
}
