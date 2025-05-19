package config

import (
	"html/template"
	"log"
)

// AppConfig holds application wide non-sensitive configuration values.
type AppConfig struct {
	InProduction  bool
	UseCache      bool
	TemplateCache map[string]*template.Template
	Port          string
	InfoLog       *log.Logger
}
