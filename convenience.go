package main

import (
	"net/http"
)

func (a *application) get(s string, h http.HandlerFunc) {
	a.App.HTTP.Router.Get(s, h)
}
