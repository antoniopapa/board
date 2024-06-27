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
	host string
}

func main() {
	//Set up logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	//Get env variables
	tier := os.Getenv("TIER")
	if len(tier) == 0 {
		tier = "starter"
	}
	host := os.Getenv("HOST")
	if len(host) == 0 {
		logger.Error("Must specify host name.")
		os.Exit(1)
	}
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
		tier: tier,
		host: host,
	}

	err = http.ListenAndServe(":8080", app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}