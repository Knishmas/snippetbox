package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.andresmacias.dev/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}
	//Retrieving and storing all page templates in a slice
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}
	//For every page we're retrieving the name of the file,
	//creating a string slice of the partial templates and the base file
	//and them parsing them into a template set. We then add them into the cache.
	for _, page := range pages {

		name := filepath.Base(page)

		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		ts, err := template.ParseFiles(files...)

		if err != nil {
			return nil, err
		}
		//Add the template set to the cache, using the name of the page
		cache[name] = ts
	}
	return cache, nil
}
