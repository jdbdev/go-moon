package loggers

import (
	"log"
	"net/http"
	"os"
	"time"
)

// RequestLogger logs request information
func RequestLogger(r *http.Request, start time.Time) {
	logger := log.New(os.Stdout, "http request: ", log.LstdFlags)
	duration := time.Since(start)
	logger.Printf("method: %s, path: %s, from: %s, duration: %v", r.Method, r.URL.Path, r.RemoteAddr, duration)
}
