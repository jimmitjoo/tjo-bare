package middleware

import (
	"context"
	"myapp/data"
	"net/http"

	"github.com/jimmitjoo/gemquick"
)

// Middleware holds application middleware with access to the framework and models
type Middleware struct {
	App    *gemquick.Gemquick
	Models data.Models
}

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

const (
	// UserContextKey is the key for storing user in request context
	UserContextKey contextKey = "user"
)

// RequireAuth is middleware that checks if a user is authenticated.
// If not authenticated, it redirects to /login.
// Usage: route.use(m.RequireAuth)
func (m *Middleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := m.App.Session.GetInt(r.Context(), "user_id")
		if userID == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RequireAuthAPI is middleware for API endpoints that returns 401 instead of redirecting.
// Usage: route.use(m.RequireAuthAPI)
func (m *Middleware) RequireAuthAPI(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := m.App.Session.GetInt(r.Context(), "user_id")
		if userID == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "unauthorized"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// LoadUser loads the current user into the request context.
// The user can be retrieved with GetUserFromContext().
func (m *Middleware) LoadUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := m.App.Session.GetInt(r.Context(), "user_id")
		if userID > 0 {
			user, err := m.Models.GetUserByID(userID)
			if err == nil {
				ctx := context.WithValue(r.Context(), UserContextKey, user)
				r = r.WithContext(ctx)
			}
		}

		next.ServeHTTP(w, r)
	})
}

// GetUserFromContext retrieves the user from the request context.
// Returns nil if no user is in context.
func GetUserFromContext(r *http.Request) *data.User {
	user, ok := r.Context().Value(UserContextKey).(*data.User)
	if !ok {
		return nil
	}
	return user
}

// RequireAdmin is middleware that checks if the authenticated user is an admin.
// It requires RequireAuth or LoadUser to be called first.
func (m *Middleware) RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetUserFromContext(r)
		if user == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add admin check logic here based on your user model
		// For example: if !user.IsAdmin { ... }

		next.ServeHTTP(w, r)
	})
}

// CORS adds Cross-Origin Resource Sharing headers for API endpoints.
// This is useful when your API is accessed from a different domain.
func (m *Middleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
