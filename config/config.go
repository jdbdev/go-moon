package config

import (
	"html/template"
	"log"
)

// AppConfig holds application runtime state and resources.
// Unlike environment config (EnvConfig), this holds:
// - Shared resources (template cache, loggers)
// - Runtime state that can change
// - Resources that need to be shared between packages
type AppConfig struct {
	// Runtime settings (initialized from environment)
	InProduction bool
	UseCache     bool
	Port         string

	// Shared resources
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger

	// Add other runtime resources here:
	// - Database connections
	// - Cache instances
	// - Session managers
	// - etc.
}
