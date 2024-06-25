package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	//Set up static file server
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//Set up routes
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /", app.notFound)
	mux.HandleFunc("GET /vpn/{$}", app.vpn)

	return mux
}