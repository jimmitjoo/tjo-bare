package handlers

import (
	"net/http"
	"time"

	"github.com/jimmitjoo/tjo"
)

// Handlers holds HTTP request handlers with access to the framework.
// Add your handlers here or generate with: tjo make handler <name>
type Handlers struct {
	App *tjo.Tjo
}

// Home renders the home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.App.HTTP.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.Logging.Error.Println("error rendering:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
