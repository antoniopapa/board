package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {
	appLinks := getAppLinks(app.tier, app.domain)

	data := templateData{ 
		AppLinks: appLinks,
		Tier: app.tier,
		BaseUri: app.domain,
	}

	app.render(writer, request, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) notFound(writer http.ResponseWriter, request *http.Request) {
	app.render(writer, request, http.StatusNotFound, "404.tmpl.html", templateData{})
}

func (app *application) vpn(writer http.ResponseWriter, request *http.Request) {
	app.render(writer, request, http.StatusOK, "vpn.tmpl.html", templateData{})
}

func (app *application) appList(writer http.ResponseWriter, request *http.Request) {
	appLinks := getAppLinks(app.tier, app.domain)
    writer.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(appLinks)

	if err != nil {
		app.serverError(writer, request, err)
		return
	}

	writer.Write(data)
}