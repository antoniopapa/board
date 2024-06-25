package main

import "net/http"

func (app *application) home(writer http.ResponseWriter, request *http.Request) {
	app.render(writer, request, http.StatusOK, "home.tmpl.html", templateData{})
}