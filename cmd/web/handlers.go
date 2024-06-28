package main

import (
	"net/http"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {
	appLinks := getAppLinks(app.tier, app.host)

	data := templateData{ 
		AppLinks: appLinks,
		Tier: app.tier,
		BaseUri: app.host,
	}

	app.render(writer, request, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) notFound(writer http.ResponseWriter, request *http.Request) {
	app.render(writer, request, http.StatusNotFound, "404.tmpl.html", templateData{})
}

func (app *application) vpn(writer http.ResponseWriter, request *http.Request) {
	app.render(writer, request, http.StatusOK, "vpn.tmpl.html", templateData{})
}