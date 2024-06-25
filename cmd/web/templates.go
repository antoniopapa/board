package main

import (
	"html/template"
	"path/filepath"
)

type templateData struct {
	BaseUri string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		//Get name of page
		name := filepath.Base(page)

		//Parse main template file
		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}
		//Parse partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}
		//Parse page
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		//set in cache
		cache[name] = ts
	}

	return cache, nil
}