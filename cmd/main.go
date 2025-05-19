package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/pkg/handlers"
	"github.com/jdbdev/go-moon/pkg/loggers"
	"github.com/jdbdev/go-moon/pkg/render"
)

// Keep main.go limited to configuration and starting/closing services.
// http.Handler and routes are in cmd/routes.go
// middleware in pkg/middleware/*
// handlers in pkg/handlers/*
// renderers in pkg/render/*

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
	app.UseCache = true
	app.InProduction = false
	app.Port = portNumber
	app.TemplateCache = tc

	// Pass app (config settings) to other packages
	render.GetConfig(&app)
	handlers.GetConfig(&app)
	loggers.GetConfig(&app)

	// Server
	srv := &http.Server{
		Addr:         portNumber,
		Handler:      routes(&app),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	loggers.ServerStartLogger()

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// DB Updater
}
