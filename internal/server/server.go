package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"geobase/internal/models"
	"geobase/internal/storage"
)

// Server is a main application server and handler structure
type Server struct {
	r            *mux.Router
	st           *storage.Storage
	srv          *http.Server
	timeout      int
	GoogleApiKey string
}

// New creates a new server
func New(st *storage.Storage, cfg *models.Config) *Server {
	srv := Server{
		r:            mux.NewRouter(),
		st:           st,
		timeout:      cfg.ReqTimeoutSec,
		GoogleApiKey: cfg.GoogleAPIKey,
	}
	srv.setupRouter()

	address := fmt.Sprintf(":%s", cfg.AppPort)
	srv.srv = &http.Server{
		Handler: srv.r,
		Addr:    address,
	}
	fmt.Println(srv.srv.Addr)

	return &srv
}

func (s *Server) setupRouter() {
	s.r.HandleFunc("/hello", s.hello).Methods("GET", "POST")
	s.r.HandleFunc("/", s.hello).Methods("GET", "POST")
	s.r.HandleFunc("/waste/type/{type_id}/location", s.getLocationURLByWasteType).
		Queries("latitude", "{latitude}", "longitude", "{longitude}").
		Methods("GET")
	s.r.HandleFunc("/waste/type/{type_id}/point", s.getLocationPointByWasteType).
		Queries("latitude", "{latitude}", "longitude", "{longitude}").
		Methods("GET")
}

// Run server
func (s *Server) Run() error {

	return s.srv.ListenAndServe()
}
