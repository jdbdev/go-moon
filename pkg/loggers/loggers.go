package loggers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jdbdev/go-moon/config"
)

// StartServer logs server start
func ServerStartLogger(app *config.AppConfig) {
	logger := log.New(os.Stdout, "http server: ", log.LstdFlags)
	logger.Printf("starting server on port %s\n", app.Port)
	logger.Printf("app in Production: %t\n", app.InProduction)
}

// RequestLogger logs request information
func RequestLogger(r *http.Request, start time.Time) {
	logger := log.New(os.Stdout, "http request: ", log.LstdFlags)
	duration := time.Since(start)
	logger.Printf("method: %s, path: %s, from: %s, duration: %v", r.Method, r.URL.Path, r.RemoteAddr, duration)
}
