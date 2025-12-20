package main

import (
	"log"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/jimmitjoo/gemquick"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Gemquick framework
	gem := &gemquick.Gemquick{}
	err = gem.New(path)
	if err != nil {
		log.Fatal(err)
	}

	gem.AppName = "myapp"

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
