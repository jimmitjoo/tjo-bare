package handlers

import (
	"net/http"
)

// render is a helper to render templates
func (h *Handlers) render(w http.ResponseWriter, r *http.Request, template string, variables, data interface{}) error {
	return h.App.HTTP.Render.Page(w, r, template, variables, data)
}
