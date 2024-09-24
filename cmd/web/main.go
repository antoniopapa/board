package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
	templateCache map[string]*template.Template
	tier string
	domain string
	appLinks []AppLink
}

func main() {
	//Set up logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	//Get env variables
	tier := os.Getenv("TIER")
	if len(tier) == 0 {
		tier = "good"
	}
	domain := os.Getenv("DOMAIN")
	if len(domain) == 0 {
		logger.Error("Must specify domain name.")
		os.Exit(1)
	}
	//Set up template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	//Set up app links cache
	appLinks := getAppLinks(tier, domain)
	//Set up application data
	app := application {
		logger: logger,
		templateCache: templateCache,
		tier: tier,
		domain: domain,
		appLinks: appLinks,
	}
	//start server
	err = http.ListenAndServe(":8080", app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}