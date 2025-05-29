package handlers

import "net/http"

//==============================================================================
// Health Check Handler
//==============================================================================

// HealthCheck responds to /healthz requests.
// Used by monitoring systems to check if service is alive.
type HealthCheck struct{}

// HealthCheck implements the Handler Interface for HealthCheck
// and writes a header with status code 204
func (h *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	message := "health check"
	w.Write([]byte(message))
}
