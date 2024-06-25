package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) render(writer http.ResponseWriter, request *http.Request, status int, page string, data templateData) {
	// Get template
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		//Should this be a 404?
		app.serverError(writer, request, err)
		return
	}

	// Write buffer
	templateBuffer := new(bytes.Buffer)

	err := ts.ExecuteTemplate(templateBuffer, "base", data)
	if err != nil {
		app.serverError(writer, request, err)
		return
	}
	// If no error, write to the response writer
	writer.WriteHeader(status)
	templateBuffer.WriteTo(writer)
}