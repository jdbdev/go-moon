package render

import (
	"fmt"
	"net/http"
	"text/template"
)

type TemplateCache struct {
}

// RenderTemplate takes two arguments; a responsewriter and a string
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
