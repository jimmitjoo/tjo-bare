package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jimmitjoo/tjo/security"
)

func (a *application) routes() *chi.Mux {
	// Security headers, CORS, request size cap and timeout. Must be registered
	// before any route. Without it an application ships with no CSP, no HSTS
	// and no frame protection, and CORS_ALLOWED_ORIGINS is read by nothing.
	a.App.HTTP.Router.Use(security.SecureMiddleware(a.securityConfig()))

	// Public routes
	a.App.HTTP.Router.Get("/", a.Handlers.Home)

	// Static files
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.HTTP.Router.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.HTTP.Router
}

// securityConfig relaxes CSP and drops HSTS in debug mode, and reads the
// allowed CORS origins from the environment. An empty CORS_ALLOWED_ORIGINS
// blocks all cross-origin requests, which is the intended default.
func (a *application) securityConfig() security.SecurityConfig {
	config := security.ProductionSecurityConfig()
	if a.App.Debug {
		config = security.DevelopmentSecurityConfig()
	}

	if origins := strings.TrimSpace(os.Getenv("CORS_ALLOWED_ORIGINS")); origins != "" {
		config.AllowedOrigins = strings.Split(origins, ",")
		for i := range config.AllowedOrigins {
			config.AllowedOrigins[i] = strings.TrimSpace(config.AllowedOrigins[i])
		}
	}

	return config
}
