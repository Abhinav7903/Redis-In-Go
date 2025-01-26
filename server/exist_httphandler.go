package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handlerExisthttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key, ok := vars["key"]
		if !ok {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}
		if s.store.Exists(key) {
			s.respond(w, ResponseMsg{Message: "success", Data: "OK"}, http.StatusOK, nil)
		} else {
			s.respond(w, ResponseMsg{Message: "error", Data: "Key does not exist"}, http.StatusNotFound, nil)
		}
	}
}
