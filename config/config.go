package config

import (
	"html/template"
	"log"
)

// AppConfig holds application wide runtime configuration and shared resources.
// This differs from pkg/config which handles environment and startup configuration.
type AppConfig struct {
	// Runtime settings
	InProduction bool
	UseCache     bool
	Port         string

	// Shared resources
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
