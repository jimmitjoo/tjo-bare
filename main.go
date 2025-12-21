package main

import (
	"log"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/jimmitjoo/tjo"
)

type application struct {
	App        *tjo.Tjo
	Handlers   *handlers.Handlers
	Middleware *middleware.Middleware
}

func main() {
	app := initApplication()
	if err := app.App.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
