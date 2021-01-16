package handler

import (
	"net/http"

	"github.com/CristianPadilla/APIs/clase3-CRUD/middleware"
)

//RoutePerson .
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)
	mux.HandleFunc("/v1/persons/create", middleware.Log(middleware.CheckAuthentication((h.create))))
	mux.HandleFunc("/v1/persons/update", middleware.Log(h.update))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/get-by-id", middleware.Log(h.getByID))
	mux.HandleFunc("/v1/persons/get-all", middleware.TimeElapsed(h.getAll))

}

//RouteLogin .
func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)
	mux.HandleFunc("/v1/login", h.login)

}
