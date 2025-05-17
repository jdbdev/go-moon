package main

import (
	"fmt"
	"log"
	"net/http"
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
	app.InProduction = false
	fmt.Printf("starting server on port %s\n", portNumber)
	fmt.Printf("app in Production: %t\n", app.InProduction)

	srv := &http.Server{
		Addr:         portNumber,
		Handler:      routes(&app),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// DB Updater
}
