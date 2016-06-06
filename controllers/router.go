package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
)

//TODO move this file somewhere else
//TODO auto register default routes (send type in param or something)
//The order is important
func RegisterRoutes(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("")))
	r.PathPrefix("/views/").Handler(http.FileServer(http.Dir("")))//todo THIS MAY NOT BE SECURE

	ru := r.PathPrefix("/user").Subrouter()
	ru.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(GetUserById)
	ru.Methods("GET").HandlerFunc(GetUser)
	ru.Methods("POST").HandlerFunc(CreateUser)
	ru.Methods("PUT").HandlerFunc(UpdateUser)
	ru.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(DeleteUser)

}