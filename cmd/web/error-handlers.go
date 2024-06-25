package main

import (
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(writer http.ResponseWriter, request *http.Request, err error) {
	var(
		method = request.Method
		uri = request.URL.RequestURI()
		trace = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}