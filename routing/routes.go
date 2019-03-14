package routing

import (
	_ "github.com/gorilla/mux"
)

func (s *Server) Routes() {
	s.Router.HandleFunc("/trip/{id}", s.GetTripHandler).Methods("GET")
	s.Router.HandleFunc("/trip/{id}", s.CreateTripHandler).Methods("POST")
	s.Router.HandleFunc("/trip/{id}", s.UpdateTripHandler).Methods("PUT")
	s.Router.HandleFunc("/trip/{id}", s.DeleteHandler).Methods("DELETE")
	s.Router.HandleFunc("/log", s.LogHandler).Methods("GET")
}
