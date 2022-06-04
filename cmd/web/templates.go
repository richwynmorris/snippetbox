package main

import (
	"html/template"
	"path/filepath"

	"richwynmorris.co.uk/internal/models"
)

// templateData is a type that acts as the holding struct for dynamic data to be passed to html templates.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Parse base template
		templateSet, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Parse all partials and add them to the template set
		templateSet, err = templateSet.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Parse the page specific to the mapCache
		templateSet, err = templateSet.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = templateSet
	}

	return cache, nil
}
