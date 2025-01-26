package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handlerDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract variables from the URL
		vars := mux.Vars(r)
		key, ok := vars["key"]
		if !ok {
			s.respond(w, ResponseMsg{Message: "error", Data: "Key is required"}, http.StatusBadRequest, nil)
			return
		}

		// Delete values from the store
		if err := s.store.Delete(key); err != nil {
			s.respond(w, ResponseMsg{Message: "error", Data: fmt.Sprintf("Error deleting key '%s': %v", key, err)}, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "success", Data: "Deleted"}, http.StatusOK, nil)
	}
}
