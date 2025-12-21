package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/jimmitjoo/tjo"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Tjo framework
	gem := &tjo.Tjo{}
	err = gem.New(path)
	if err != nil {
		log.Fatal(err)
	}

	gem.AppName = "myapp"

	// Initialize upper/db session if database is configured
	if gem.Data.DB.Pool != nil {
		if err := data.InitDB(gem.Data.DB.Pool); err != nil {
			log.Printf("Warning: Could not initialize upper/db: %v", err)
		}
	}

	myMiddleware := &middleware.Middleware{
		App: gem,
	}

	myHandlers := &handlers.Handlers{
		App: gem,
	}

	app := &application{
		App:        gem,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.HTTP.Router = app.routes()

	return app
}
