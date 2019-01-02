package routing

import (
	"bitbucket.org/big_life/big-life-backend/trip/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	tripService service.TripServiceInt
	Router      *mux.Router
}

func (s Server) Start() {
	s.Routes()
	http.ListenAndServe(":8080", s.Router)
}

func (s *Server) Routes() {
	s.Router.HandleFunc("/trip/{id}", s.GetTripHandler).Methods("GET")
	s.Router.HandleFunc("/trip/{id}", s.CreateTripHandler).Methods("POST")
	s.Router.HandleFunc("/trip/{id}", s.UpdateTripHandler).Methods("PUT")
	s.Router.HandleFunc("/trip/{id}", s.DeleteHandler).Methods("DELETE")
}

func NewServer(ts service.TripServiceInt, mux *mux.Router) *Server {
	return &Server{tripService: ts, Router: mux}
}
