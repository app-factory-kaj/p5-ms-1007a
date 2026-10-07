// Package httpapi implements the greeter service's HTTP handlers and the
// shared JSON response conventions from app-factory-kaj/e2e-reference.
package httpapi

import (
	"encoding/json"
	"net/http"
)

// NewRouter wires every route the greeter service exposes.
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handleHealth)
	return mux
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
