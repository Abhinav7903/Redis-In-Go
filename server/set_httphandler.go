package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// handlerSet returns an HTTP handler for setting values for a key.
func (s *Server) handlerSet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Extract variables from the URL
		key, ok := vars["key"]
		if !ok {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		// Parse the request body for values
		var values []string
		if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
			http.Error(w, "Invalid request body. Expected a JSON array of values.", http.StatusBadRequest)
			return
		}

		// Set values in the store
		if err := s.store.Set(key, values...); err != nil {
			http.Error(w, fmt.Sprintf("Error setting values for key '%s': %v", key, err), http.StatusInternalServerError)
			return
		}

		s.respond(w, ResponseMsg{Message: "success", Data: "OK "}, http.StatusOK, nil)
	}
}

// handlerSetUnique returns an HTTP handler for setting unique values for a key.
func (s *Server) handlerSetUnique() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Extract variables from the URL
		key, ok := vars["key"]
		if !ok {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		// Parse the request body for values
		var values []string
		if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
			http.Error(w, "Invalid request body. Expected a JSON array of values.", http.StatusBadRequest)
			return
		}

		// Set unique values in the store
		if err := s.store.SetUnique(key, values...); err != nil {
			http.Error(w, fmt.Sprintf("Error setting unique values for key '%s': %v", key, err), http.StatusInternalServerError)
			return
		}

		// Respond with a structured success message
		s.respond(w, ResponseMsg{Message: "success", Data: "OK"}, http.StatusOK, nil)
	}
}
