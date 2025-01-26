package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// handlerGet returns an HTTP handler for retrieving values based on a key.
func (s *Server) handlerGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Extract variables from the URL
		key, ok := vars["key"]
		if !ok {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		// Fetch values from the store
		values, err := s.store.Get(key)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving key '%s': %v", key, err), http.StatusInternalServerError)
			return
		}

		// Handle case where no values are found
		if len(values) == 0 {
			http.Error(w, fmt.Sprintf("No values found for key: %s", key), http.StatusNotFound)
			return
		}

		// Respond with JSON containing the values
		response := map[string]interface{}{
			"key":    key,
			"values": values,
		}

		s.respond(w, ResponseMsg{Message: "success", Data: response}, http.StatusOK, nil)

	}
}

// handlerGetUnique returns an HTTP handler for retrieving unique values based on a key.
func (s *Server) handlerGetUnique() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Extract variables from the URL
		key, ok := vars["key"]
		if !ok {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		// Fetch unique values from the store
		values, err := s.store.GetUnique(key)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving unique values for key '%s': %v", key, err), http.StatusInternalServerError)
			return
		}

		// Handle case where no unique values are found
		if len(values) == 0 {
			http.Error(w, fmt.Sprintf("No unique values found for key: %s", key), http.StatusNotFound)
			return
		}

		// Respond with JSON containing the unique values
		response := map[string]interface{}{
			"key":    key,
			"values": values,
		}
		s.respond(w, ResponseMsg{Message: "success", Data: response}, http.StatusOK, nil)
	}
}

// handlerGetKey returns an HTTP handler for retrieving keys based on a value.
func (s *Server) handlerGetKey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Extract variables from the URL
		value, ok := vars["value"]
		if !ok {
			http.Error(w, "Value is required", http.StatusBadRequest)
			return
		}

		// Fetch keys from the store
		keys, err := s.store.GetKeyFromValue(value)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving keys for value '%s': %v", value, err), http.StatusInternalServerError)
			return
		}

		// Handle case where no keys are found
		if len(keys) == 0 {
			http.Error(w, fmt.Sprintf("No keys found for value: %s", value), http.StatusNotFound)
			return
		}

		// Respond with JSON containing the keys
		response := map[string]interface{}{
			"value": value,
			"keys":  keys,
		}

		s.respond(w, ResponseMsg{Message: "success", Data: response}, http.StatusOK, nil)

	}
}
