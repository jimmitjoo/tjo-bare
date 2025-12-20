package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (route *application) routes() *chi.Mux {
	// Middleware must come before any routes

	// Load user into context for all routes
	route.use(route.Middleware.LoadUser)

	// Public routes
	route.get("/", route.Handlers.Home)

	// API routes (exempt from CSRF protection by framework)
	route.App.Routes.Route("/api", func(r chi.Router) {
		// Apply CORS for API routes
		r.Use(route.Middleware.CORS)

		// Public API endpoints
		r.Get("/health", route.Handlers.APIHealthCheck)
		r.Post("/login", route.Handlers.APILogin)
		r.Post("/logout", route.Handlers.APILogout)

		// Protected API endpoints
		r.Group(func(r chi.Router) {
			r.Use(route.Middleware.RequireAuthAPI)

			// User CRUD
			r.Get("/users", route.Handlers.APIGetUsers)
			r.Get("/users/{id}", route.Handlers.APIGetUser)
			r.Post("/users", route.Handlers.APICreateUser)
			r.Delete("/users/{id}", route.Handlers.APIDeleteUser)
		})
	})

	// Protected web routes example
	route.App.Routes.Group(func(r chi.Router) {
		r.Use(route.Middleware.RequireAuth)

		// Add protected routes here
		// r.Get("/dashboard", route.Handlers.Dashboard)
	})

	// Static routes for images and other static files
	fileServer := http.FileServer(http.Dir("./public"))
	route.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return route.App.Routes
}
