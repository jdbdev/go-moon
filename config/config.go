package config

import (
	"html/template"
	"log"

	"github.com/jdbdev/go-moon/pkg/env"
)

// AppConfig holds application runtime state and resources.
// It combines:
// 1. Runtime settings (from environment, but can change during runtime)
// 2. Shared resources (created and shared between packages)
type AppConfig struct {
	// Runtime settings that can change during application lifecycle
	// Initially set from env.ServerConfig but may be modified
	Runtime RuntimeConfig

	// Shared resources that are created after startup
	Resources ResourceConfig
}

// RuntimeConfig holds settings that can change during runtime
// Initially populated from env.ServerConfig
type RuntimeConfig struct {
	env.ServerConfig
}

// ResourceConfig holds shared resources created after startup
type ResourceConfig struct {
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}

// NewAppConfig creates a new AppConfig initialized from environment settings
func NewAppConfig(envCfg env.EnvConfig) *AppConfig {
	return &AppConfig{
		Runtime: RuntimeConfig{
			ServerConfig: envCfg.Server,
		},
		Resources: ResourceConfig{
			TemplateCache: make(map[string]*template.Template),
			InfoLog:       log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime),
		},
	}
}
