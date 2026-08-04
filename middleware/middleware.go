package middleware

import (
	"myapp/data"

	"github.com/jimmitjoo/tjo"
)

// Middleware holds application middleware with access to the framework.
// Add your custom middleware here or generate with: tjo make middleware <name>
type Middleware struct {
	App    *tjo.Tjo
	Models data.Models
}

// Example middleware template:
//
// func (m *Middleware) MyMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Before request
//         next.ServeHTTP(w, r)
//         // After request
//     })
// }
