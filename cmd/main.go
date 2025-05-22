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

func main() {
	//==========================================================================
	// Configuration & Setup
	//==========================================================================

	// Load environment configuration
	envCfg := env.LoadEnv()

	// Initialize application config
	app := config.NewAppConfig(envCfg)

	// Initialize template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.Resources.TemplateCache = tc

	// Initialize logger
	logger := loggers.NewLogger(loggers.LoggerConfig{
		Port:         app.Runtime.Port,
		InProduction: app.Runtime.InProduction,
		UseCache:     app.Runtime.UseCache,
	})

	// Log initial configuration
	logger.LogConfig()

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

	//==========================================================================
	// Server Configuration & Startup
	//==========================================================================

	// Initialize HTTP server
	srv := &http.Server{
		Addr:         ":" + app.Runtime.Port,
		Handler:      routes(app, logger), // Pass app config and logger
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Start server in background
	go func() {
		logger.LogServerStart()
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
	logger.LogServerStart() // Log shutdown start

	// Shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
