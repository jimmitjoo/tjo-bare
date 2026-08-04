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
	app := &tjo.Tjo{}
	err = app.New(path)
	if err != nil {
		log.Fatal(err)
	}

	app.AppName = "myapp"

	// Initialize upper/db session and models if a database is configured.
	// Without one, models stay zero-valued and any model call will fail.
	var models data.Models
	if app.Data.DB.Pool != nil {
		if err := data.InitDB(app.Data.DB.Pool); err != nil {
			log.Printf("Warning: Could not initialize upper/db: %v", err)
		}
		models, err = data.New(app.Data.DB.Pool)
		if err != nil {
			log.Fatal(err)
		}
	}

	myMiddleware := &middleware.Middleware{
		App:    app,
		Models: models,
	}

	myHandlers := &handlers.Handlers{
		App:    app,
		Models: models,
	}

	a := &application{
		App:        app,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	a.App.HTTP.Router = a.routes()

	return a
}
