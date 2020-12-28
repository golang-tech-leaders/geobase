package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"geobase/internal/config"
	"geobase/internal/logger"
)

// Server provides the server functionality
type Server struct {
	r   *mux.Router
	srv *http.Server
	log *logger.Logger

	urlFinder URLFinder
	locFinder LocationFinder
}

// NewServer creates a server and prepares a router
func NewServer(cfg *config.AppConfig,
	urlFinder URLFinder, locFinder LocationFinder, logger *logger.Logger) *Server {
	s := Server{
		r:         mux.NewRouter(),
		log:       logger,
		urlFinder: urlFinder,
		locFinder: locFinder,
	}

	s.setupRouter()

	address := fmt.Sprintf(":%s", cfg.AppPort)
	s.srv = &http.Server{
		Handler:      s.r,
		Addr:         address,
		WriteTimeout: time.Duration(cfg.ReqTimeoutSec) * time.Second,
	}

	return &s
}

func (s *Server) setupRouter() {
	s.r.HandleFunc("/waste/type/{type_id}/location", s.getLocURLForWasteType).Methods("GET")
	s.r.HandleFunc("/waste/type/{type_id}/point", s.getLocPointForWasteType).Methods("GET")
	s.r.HandleFunc("/waste/type/{type_id}/points", s.getLocPointListForWasteType).Methods("GET")

}

// Run starts the server
func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

// Shutdown closes server
func (s *Server) Shutdown() error {
	return s.srv.Close()
}
