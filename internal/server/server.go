package server

import (
	"fmt"
	"geobase/internal/config"
	"geobase/internal/logger"
	"net/http"

	"github.com/gorilla/mux"
)

// Server provides the server functionality
type Server struct {
	r       *mux.Router
	db      GeobaseRepository
	srv     *http.Server
	timeout int
	log     *logger.Logger
}

// NewServer creates a server and prepares a router
func NewServer(cfg *config.AppConfig, storage GeobaseRepository, logger *logger.Logger) *Server {
	s := Server{
		r:       mux.NewRouter(),
		db:      storage,
		timeout: cfg.ReqTimeoutSec,
		log:     logger,
	}

	s.setupRouter()

	address := fmt.Sprintf(":%s", cfg.AppPort)
	s.srv = &http.Server{
		Handler: s.r,
		Addr:    address,
	}

	return &s
}

func (s *Server) setupRouter() {
	s.r.HandleFunc("/waste/type/{type_id}/location", s.getLocForWasteType).Methods("GET")
}

// Run starts the server
func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

// Shutdown closes server
func (s *Server) Shutdown() error {
	return s.srv.Close()
}
