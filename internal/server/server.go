package server

import (
	"fmt"
	"net/http"

	"Go-Boilerplate/internal/config"
	"Go-Boilerplate/internal/handler"
	"Go-Boilerplate/internal/middleware"
)

// Server configuration.
type Server struct {
	config     *config.Config
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	// Create a new ServeMux and register routes.
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handler.LoginCheckHandler)

	// Wrap the mux with logging middleware.
	loggedMux := middleware.Logger(mux)

	return &Server{
		config: cfg,
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: loggedMux,
		},
	}
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	fmt.Printf("Server starting on port %s\n", s.config.Port)
	return s.httpServer.ListenAndServe()
}
