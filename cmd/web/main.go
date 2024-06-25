package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"flag"
)

//TODO Get images for each app
//TODO Stylesheet
//TODO VPN Page
//TODO Auth through LDAP for VPN page

type application struct {
	logger *slog.Logger
	templateCache map[string]*template.Template
	tier string
}

func main() {
	//Get env variables
	addr := flag.String("addr", ":4000", "HTTP network address")
	tier := flag.String("tier", "starter", "tier the core is on")
	flag.Parse()
	//Set up logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	//Set up template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	//Set up application data
	app := application {
		logger: logger,
		templateCache: templateCache,
		tier: *tier,
	}

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}