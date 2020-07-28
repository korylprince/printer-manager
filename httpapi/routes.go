package httpapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/korylprince/httputil/jsonapi"
	"github.com/korylprince/printer-manager/db/crud"
)

const uuidRegexp = "(?i)[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}"

//API is the current API version
const API = "1.0"
const apiPath = "/api/" + API

//Router returns a new API router
func (s *Server) Router() http.Handler {
	r := mux.NewRouter()

	apirouter := jsonapi.New(s.output, s.auth, s.sessionStore, nil)
	r.PathPrefix(apiPath).Handler(http.StripPrefix(apiPath, apirouter))

	apirouter.Handle("GET", "/sync/stats",
		s.syncStats, true)
	apirouter.Handle("POST", "/sync/trigger",
		s.syncTrigger, true)

	apirouter.HandleTX("POST", "/buildings",
		s.db, crud.CreateBuilding, true)
	apirouter.HandleTX("GET", fmt.Sprintf("/buildings/{id:%s}", uuidRegexp),
		s.db, crud.ReadBuilding, true)
	apirouter.HandleTX("PUT", fmt.Sprintf("/buildings/{id:%s}", uuidRegexp),
		s.db, crud.UpdateBuilding, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf("/buildings/{id:%s}", uuidRegexp),
		s.db, crud.DeleteBuilding, true)
	apirouter.HandleTX("GET", "/buildings",
		s.db, crud.QueryBuildings, true)

	apirouter.HandleTX("POST", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations", uuidRegexp),
		s.db, crud.CreateLocation, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{id:%s}", uuidRegexp, uuidRegexp),
		s.db, crud.ReadLocation, true)
	apirouter.HandleTX("PUT", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{id:%s}", uuidRegexp, uuidRegexp),
		s.db, crud.UpdateLocation, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{id:%s}", uuidRegexp, uuidRegexp),
		s.db, crud.DeleteLocation, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations", uuidRegexp),
		s.db, crud.QueryLocations, true)
	apirouter.HandleTX("GET", "/locations",
		s.db, crud.QueryLocations, true)

	apirouter.HandleTX("POST", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/printers",
		uuidRegexp, uuidRegexp),
		s.db, crud.CreatePrinter, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/printers/{id:%s}",
		uuidRegexp, uuidRegexp, uuidRegexp),
		s.db, crud.ReadPrinter, true)
	apirouter.HandleTX("PUT", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/printers/{id:%s}",
		uuidRegexp, uuidRegexp, uuidRegexp),
		s.db, crud.UpdatePrinter, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/printers/{id:%s}",
		uuidRegexp, uuidRegexp, uuidRegexp),
		s.db, crud.DeletePrinter, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/printers",
		uuidRegexp, uuidRegexp),
		s.db, crud.QueryPrinters, true)
	apirouter.HandleTX("GET", "/printers",
		s.db, crud.QueryPrinters, true)

	apirouter.HandleTX("POST", "/manufacturers",
		s.db, crud.CreateManufacturer, true)
	apirouter.HandleTX("GET", fmt.Sprintf("/manufacturers/{id:%s}", uuidRegexp),
		s.db, crud.ReadManufacturer, true)
	apirouter.HandleTX("PUT", fmt.Sprintf("/manufacturers/{id:%s}", uuidRegexp),
		s.db, crud.UpdateManufacturer, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf("/manufacturers/{id:%s}", uuidRegexp),
		s.db, crud.DeleteManufacturer, true)
	apirouter.HandleTX("GET", "/manufacturers",
		s.db, crud.QueryManufacturers, true)

	apirouter.HandleTX("POST", fmt.Sprintf(
		"/manufacturers/{manufacturer_id:%s}/models", uuidRegexp),
		s.db, crud.CreateModel, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/manufacturers/{manufacturer_id:%s}/models/{id:%s}", uuidRegexp, uuidRegexp),
		s.db, crud.ReadModel, true)
	apirouter.HandleTX("PUT", fmt.Sprintf(
		"/manufacturers/{manufacturer_id:%s}/models/{id:%s}", uuidRegexp, uuidRegexp),
		s.db, crud.UpdateModel, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf(
		"/manufacturers/{manufacturer_id:%s}/models/{id:%s}", uuidRegexp, uuidRegexp),
		s.db, crud.DeleteModel, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/manufacturers/{manufacturer_id:%s}/models", uuidRegexp),
		s.db, crud.QueryModels, true)
	apirouter.HandleTX("GET", "/models",
		s.db, crud.QueryModels, true)

	apirouter.HandleTX("POST", "/users",
		s.db, crud.CreateUser, true)
	apirouter.HandleTX("GET", fmt.Sprintf("/users/{id:%s}", uuidRegexp),
		s.db, crud.ReadUser, true)
	apirouter.HandleTX("PUT", fmt.Sprintf("/users/{id:%s}", uuidRegexp),
		s.db, crud.UpdateUser, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf("/users/{id:%s}", uuidRegexp),
		s.db, crud.DeleteUser, true)
	apirouter.HandleTX("GET", "/users",
		s.db, crud.QueryUsers, true)

	apirouter.HandleTX("POST", "/groups",
		s.db, crud.CreateGroup, true)
	apirouter.HandleTX("GET", fmt.Sprintf("/groups/{id:%s}", uuidRegexp),
		s.db, crud.ReadGroup, true)
	apirouter.HandleTX("PUT", fmt.Sprintf("/groups/{id:%s}", uuidRegexp),
		s.db, crud.UpdateGroup, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf("/groups/{id:%s}", uuidRegexp),
		s.db, crud.DeleteGroup, true)
	apirouter.HandleTX("GET", "/groups",
		s.db, crud.QueryGroups, true)

	apirouter.HandleTX("GET", fmt.Sprintf(
		"/groups/{group_id:%s}/users",
		uuidRegexp),
		s.db, crud.ReadGroupUsers, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/users/{user_id:%s}/groups",
		uuidRegexp),
		s.db, crud.ReadUserGroups, true)
	apirouter.HandleTX("PUT", fmt.Sprintf(
		"/groups/{group_id:%s}/users/{user_id:%s}/assign",
		uuidRegexp, uuidRegexp),
		s.db, crud.RelateGroupUser, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf(
		"/groups/{group_id:%s}/users/{user_id:%s}/assign",
		uuidRegexp, uuidRegexp),
		s.db, crud.UnrelateGroupUser, true)

	apirouter.HandleTX("GET", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/users",
		uuidRegexp, uuidRegexp),
		s.db, crud.ReadLocationUsers, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/users/{user_id:%s}/locations",
		uuidRegexp),
		s.db, crud.ReadUserLocations, true)
	apirouter.HandleTX("PUT", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/users/{user_id:%s}/assign",
		uuidRegexp, uuidRegexp, uuidRegexp),
		s.db, crud.RelateLocationUser, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/users/{user_id:%s}/assign",
		uuidRegexp, uuidRegexp, uuidRegexp),
		s.db, crud.UnrelateLocationUser, true)

	apirouter.HandleTX("GET", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/groups",
		uuidRegexp, uuidRegexp),
		s.db, crud.ReadLocationGroups, true)
	apirouter.HandleTX("GET", fmt.Sprintf(
		"/groups/{group_id:%s}/locations",
		uuidRegexp),
		s.db, crud.ReadGroupLocations, true)
	apirouter.HandleTX("PUT", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/groups/{group_id:%s}/assign",
		uuidRegexp, uuidRegexp, uuidRegexp),
		s.db, crud.RelateLocationGroup, true)
	apirouter.HandleTX("DELETE", fmt.Sprintf(
		"/buildings/{building_id:%s}/locations/{location_id:%s}/groups/{group_id:%s}/assign",
		uuidRegexp, uuidRegexp, uuidRegexp),
		s.db, crud.UnrelateLocationGroup, true)

	apirouter.HandleTX("GET", "/users/{username}/printers",
		s.db, readUserPrinters, false)

	return r
}
