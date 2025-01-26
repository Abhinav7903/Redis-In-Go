package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ResponseMsg struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (s *Server) RegisterAPIs() {
	//  Register your API routes
	s.router.HandleFunc("/get/{key}", s.handlerGet()).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/getuq/{key}", s.handlerGetUnique()).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/getkey/{value}", s.handlerGetKey()).Methods(http.MethodGet, http.MethodOptions)

	// Set the key with values
	s.router.HandleFunc("/set/{key}", s.handlerSet()).Methods(http.MethodPost, http.MethodOptions)
	// Set the key with unique values
	s.router.HandleFunc("/setuq/{key}", s.handlerSetUnique()).Methods(http.MethodPost, http.MethodOptions)

	// DELETE the key
	s.router.HandleFunc("/delete/{key}", s.handlerDelete()).Methods(http.MethodDelete, http.MethodOptions)

	// Check if the key exists
	s.router.HandleFunc("/exists/{key}", s.handlerExisthttp()).Methods(http.MethodGet, http.MethodOptions)

	// Get the expiration time for the key
	s.router.HandleFunc("/ttl/{key}", s.handlerTTL()).Methods(http.MethodGet, http.MethodOptions)

	// Set the expiration time for the key
	s.router.HandleFunc("/expire/{key}", s.handlerExpire()).Methods(http.MethodPost, http.MethodOptions)

	// help
	s.router.HandleFunc("/help", s.handlerHelp()).Methods(http.MethodGet, http.MethodOptions)
}

func (s *Server) respond(
	w http.ResponseWriter,
	data interface{},
	status int,
	err error,
) {
	// Set content type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var resp *ResponseMsg
	if err == nil {
		resp = &ResponseMsg{
			Message: "success",
			Data:    data,
		}
	} else {
		resp = &ResponseMsg{
			Message: err.Error(),
			Data:    nil, // Ensure no conflicting message structure
		}
	}

	// Encode the response
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("Failed to encode response:", "error", err)
	}
}
