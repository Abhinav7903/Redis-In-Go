package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// handlerExpire returns an HTTP handler for setting an expiration time (TTL) for a key.
func (s *Server) handlerExpire() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Extract variables from the URL
		key, ok := vars["key"]
		if !ok {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		// Parse TTL from the query parameters
		ttlStr := r.URL.Query().Get("ttl")
		if ttlStr == "" {
			http.Error(w, "TTL parameter is required", http.StatusBadRequest)
			return
		}

		ttl, err := time.ParseDuration(ttlStr + "s")
		if err != nil {
			http.Error(w, "Invalid TTL value", http.StatusBadRequest)
			return
		}

		// Set the expiration for the key in the store
		if err := s.store.Expire(key, ttl); err != nil {
			http.Error(w, fmt.Sprintf("Error setting expiration for key '%s': %v", key, err), http.StatusInternalServerError)
			return
		}

		// Respond with success message
		s.respond(w, ResponseMsg{Message: "success", Data: "OK"}, http.StatusOK, nil)
	}
}

// handlerTTL returns an HTTP handler for retrieving the TTL (expiration time) of a key.
func (s *Server) handlerTTL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Extract variables from the URL
		key, ok := vars["key"]
		if !ok {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		// Get TTL from the store
		ttl, err := s.store.TTL(key)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving TTL for key '%s': %v", key, err), http.StatusInternalServerError)
			return
		}

		// Respond with TTL value
		s.respond(w, ResponseMsg{Message: "success", Data: fmt.Sprintf("TTL: %d seconds", int(ttl.Seconds()))}, http.StatusOK, nil)
	}
}
