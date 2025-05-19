package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/pkg/loggers"
	"github.com/jdbdev/go-moon/pkg/render"
)

// Keep main.go limited to starting and closing services.
// http.Handler and routes are in cmd/routes.go
// middleware in cmd/middleware.go
// handlers in internal/handlers/*
// renderers in internal/render/*

var app config.AppConfig

// Constants & Variables
const portNumber = ":8080"

func main() {

	// Template Cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Settings
	app.TemplateCache = tc
	app.UseCache = true
	app.InProduction = false
	app.Port = portNumber

	loggers.ConfigLogger(&app)

	// Server
	srv := &http.Server{
		Addr:         portNumber,
		Handler:      routes(&app),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	loggers.ServerStartLogger(&app)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// DB Updater
}
