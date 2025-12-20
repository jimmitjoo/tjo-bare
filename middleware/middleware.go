package middleware

import (
	"github.com/jimmitjoo/gemquick"
)

// Middleware holds application middleware with access to the framework.
// Add your custom middleware here or generate with: gq make middleware <name>
type Middleware struct {
	App *gemquick.Gemquick
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
