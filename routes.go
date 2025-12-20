package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// Public routes
	a.get("/", a.Handlers.Home)

	// Static files
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.HTTP.Router.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.HTTP.Router
}
