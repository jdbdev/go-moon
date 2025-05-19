package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jdbdev/go-moon/config"
)

var app config.AppConfig

// RenderTemplate takes two arguments; a responsewriter and a string
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// 1. Get or create a template cache
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
		fmt.Println("UseCache: true")
	} else {
		templateCache, _ = CreateTemplateCache()
		fmt.Println("UseCache: false")
	}

	// 2. Check to see if tmpl matches any index in templateCache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get the template from tempate cache")
	}

	// 3. Render template
	err := template.Execute(w, template)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTemplateCache parses files in ./templates and creates a cache for future loading.
// This avoids having to add templates manually and all templates are added when called.
// Function parses for all .tmpl files, including associated layouts
func CreateTemplateCache() (map[string]*template.Template, error) {

	// 1. Create template cache
	newCache := map[string]*template.Template{}

	// 2. Get all the files that end with *.page.tmpl from ./templates into a []string
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return newCache, err
	}

	// 3. Range through the pages (type []string)
	for _, page := range pages {
		name := filepath.Base(page)

		// templateSet includes the page template and associated layouts
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return newCache, err
		}

		layout, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return newCache, err
		}

		if len(layout) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return newCache, err
			}
		}

		newCache[name] = templateSet
	}

	return newCache, nil
}
