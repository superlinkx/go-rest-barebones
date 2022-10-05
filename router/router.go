package router

import (
	"github.com/gorilla/mux"
	"github.com/superlinkx/go-rest-barebones/app"
	"github.com/superlinkx/go-rest-barebones/controller"
)

// NewRouter creates a new gorilla/mux router with our endpoints
func NewRouter(app app.App) *mux.Router {
	customerController := controller.NewCustomerController(app)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/customer", customerController.CustomerPostHandler).Methods("POST")
	r.HandleFunc("/customer/{id:[0-9]+}", customerController.CustomerGetHandler).Methods("GET")
	r.HandleFunc("/customer/{id:[0-9]+}", customerController.CustomerPutHandler).Methods("PUT")
	r.HandleFunc("/customer/{id:[0-9]+}", customerController.CustomerDeleteHandler).Methods("DELETE")
	return r
}
