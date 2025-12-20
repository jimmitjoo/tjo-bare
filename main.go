package main

import (
	"myapp/handlers"
	"myapp/middleware"

	"github.com/jimmitjoo/gemquick"
)

type application struct {
	App        *gemquick.Gemquick
	Handlers   *handlers.Handlers
	Middleware *middleware.Middleware
}

func main() {
	app := initApplication()
	app.App.ListenAndServe()
}
