package main

import (
	"net/http"
	"strings"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {
	domain := request.Host
	domainPieces := strings.Split(domain, ".")
	var baseURI string
	if len(domainPieces) > 1 {
		baseURI = domainPieces[len(domainPieces) - 2] + domainPieces[len(domainPieces) - 1]
	} else {
		baseURI = domain
	}

	data := templateData{ 
		AppLinks: getAppLinks(app.tier, baseURI),
		Tier: app.tier,
		BaseUri: baseURI,
	}

	app.render(writer, request, http.StatusOK, "home.tmpl.html", data)
}