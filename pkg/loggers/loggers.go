package loggers

import (
	"log"
	"net/http"
	"os"
	"time"
)

// LoggerConfig holds the configuration needed for logging
type LoggerConfig struct {
	Port         string
	InProduction bool
	UseCache     bool
}

// Logger handles all application logging needs.
// It contains any configuration needed for logging
// and provides methods for different types of logs.
type Logger struct {
	LoggerConfig  // Embed the config
	serverLogger  *log.Logger
	configLogger  *log.Logger
	requestLogger *log.Logger
}

// NewLogger creates a new Logger instance with the provided configuration
func NewLogger(cfg LoggerConfig) *Logger {
	return &Logger{
		LoggerConfig:  cfg,
		serverLogger:  log.New(os.Stdout, "http server: ", log.LstdFlags),
		configLogger:  log.New(os.Stdout, "config: ", log.LstdFlags),
		requestLogger: log.New(os.Stdout, "http request: ", log.LstdFlags),
	}
}

// LogServerStart logs server startup information
func (l *Logger) LogServerStart() {
	l.serverLogger.Printf("starting server on port %s\n", l.Port)
	l.serverLogger.Printf("app in Production: %t\n", l.InProduction)
}

// LogConfig logs configuration settings at application start
func (l *Logger) LogConfig() {
	l.configLogger.Printf("InProduction: %t, UseCache: %t", l.InProduction, l.UseCache)
}

// LogRequest logs information about an HTTP request
func (l *Logger) LogRequest(r *http.Request, start time.Time) {
	duration := time.Since(start)
	l.requestLogger.Printf("method: %s, path: %s, from: %s, duration: %v",
		r.Method, r.URL.Path, r.RemoteAddr, duration)
}
