package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jdbdev/go-moon/config"
)

// Keep main.go limited to starting and closing services.
// http.Handler and routes are in cmd/routes.go
// handlers are in internal/handlers/*
// renderers are in internal/render/*

var app config.AppConfig

// Constants & Variables
const portNumber = ":8080"

func main() {

	// Server
	logger := log.New(os.Stdout, "http server: ", log.LstdFlags)
	app.InProduction = false
	logger.Printf("starting server on port %s\n", portNumber)
	logger.Printf("app in Production: %t\n", app.InProduction)

	srv := &http.Server{
		Addr:         portNumber,
		Handler:      routes(&app),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}

	// DB Updater
}
