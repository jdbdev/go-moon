package loggers

import (
	"log"
	"net/http"
	"os"
)

// Logger creates a new logger and logs request information
func Logger(r *http.Request) {
	logger := log.New(os.Stdout, "http request: ", log.LstdFlags)
	logger.Printf("method: %s, path: %s, from: %s,", r.Method, r.URL.Path, r.RemoteAddr)
}
