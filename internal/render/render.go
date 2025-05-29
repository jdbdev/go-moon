package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/jdbdev/go-moon/config"
)

// TemplateRenderer handles template rendering with configuration (conforms to Renderer interface)
type TemplateRenderer struct {
	app *config.AppConfig
}

// NewTemplateRenderer creates a new template renderer with the given config
func NewTemplateRenderer(a *config.AppConfig) *TemplateRenderer {
	return &TemplateRenderer{
		app: a,
	}
}

// RenderTemplate renders the specified template with optional data
func (tr *TemplateRenderer) RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	// 1. Get or create a template cache
	var templateCache map[string]*template.Template
	var err error

	if tr.app.Runtime.UseCache {
		templateCache = tr.app.Resources.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()
		if err != nil {
			http.Error(w, "Error creating template cache", http.StatusInternalServerError)
			return err
		}
	}

	// 2. Check to see if tmpl matches any index in templateCache
	template, ok := templateCache[tmpl]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return fmt.Errorf("template %s not found in cache", tmpl)
	}

	// 3. Render template with data
	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return err
	}

	return nil
}

// CreateTemplateCache parses files in web/templates and creates a cache for future loading.
// This avoids having to add templates manually and all templates are added when called.
// Function parses for all .tmpl files, including associated layouts
func CreateTemplateCache() (map[string]*template.Template, error) {
	// 1. Create template cache
	newCache := map[string]*template.Template{}

	// 2. Get all the files that end with *.page.tmpl from web/templates into a []string
	pages, err := filepath.Glob("./web/templates/*.page.tmpl")
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

		layout, err := filepath.Glob("./web/templates/*.layout.tmpl")
		if err != nil {
			return newCache, err
		}

		if len(layout) > 0 {
			templateSet, err = templateSet.ParseGlob("./web/templates/*.layout.tmpl")
			if err != nil {
				return newCache, err
			}
		}

		newCache[name] = templateSet
	}

	return newCache, nil
}
