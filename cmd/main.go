package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/internal/render"
	"github.com/jdbdev/go-moon/pkg/env"
	"github.com/jdbdev/go-moon/pkg/loggers"
)

//==============================================================================
// Application Entry Point
//==============================================================================

// Initialize app config variable.
// Pass by reference to other packages: &app
// From other packages, use pointer to instance: var app *config.AppConfig
var app config.AppConfig

func main() {
	//==========================================================================
	// Configuration & Setup
	//==========================================================================

	// Load environment configuration
	cfg := env.LoadEnv()

	// Initialize template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Configure application settings
	app.UseCache = cfg.Server.UseCache
	app.InProduction = cfg.Server.InProduction
	app.Port = cfg.Server.Port
	app.TemplateCache = tc

	// // Initialize database connection
	// dbConfig := database.Config{
	// 	Host:     cfg.Database.Host,
	// 	Port:     cfg.Database.Port,
	// 	User:     cfg.Database.User,
	// 	Password: cfg.Database.Password,
	// 	DBName:   cfg.Database.DBName,
	// }

	// err = database.Connect(dbConfig)
	// if err != nil {
	// 	log.Fatal("Cannot connect to database:", err)
	// }
	// defer database.Close()

	// Share configuration with packages
	loggers.GetConfig(&app)

	//==========================================================================
	// Server Configuration & Startup
	//==========================================================================

	// Initialize HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      routes(&app),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Start server in background
	go func() {
		loggers.ServerStartLogger()
		// ListenAndServe always returns an error when the server stops
		// http.ErrServerClosed is the expected error when Shutdown is called
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	//==========================================================================
	// Graceful Shutdown Handling
	//==========================================================================

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)

	// Tell the OS to send SIGINT (Ctrl+C) and SIGTERM signals to our quit channel
	// This is like registering "if you get these signals, put them in my channel"
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Wait here until we receive a signal
	// The <- operator means "receive from channel"
	// This line blocks (waits) until a signal is received
	<-quit

	// If we get here, it means we received a shutdown signal
	log.Println("Shutting down server...")

	// Shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
