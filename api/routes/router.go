package routes

import (
	"freezebee/api/controllers"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
	"freezebee/api/middleware_api"
)

// Crée un nouveau routeur en utilisant le package mux.
func NewRouter() *mux.Router {
    r := mux.NewRouter().StrictSlash(true)

    // Ajoute des middlewares globaux.
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

	r.Use(middleware_api.Middleware)

    // Charge la gestion des en-têtes CORS.
    LoadCors(r)

    // Configure les routes pour diverses actions avec leurs contrôleurs associés.

	r.HandleFunc("/freezebee", controllers.Get).Methods("GET")
	r.HandleFunc("/freezebee", controllers.Post).Methods("POST")
	r.HandleFunc("/freezebee", controllers.Patch).Methods("PATCH")
	r.HandleFunc("/freezebee", controllers.Delete).Methods("DELETE")

	return r
}