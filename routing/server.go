package routing

import (
	"net/http"
	"big-life-backend/trip/service"

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

func NewServer(ts service.TripServiceInt, mux *mux.Router) *Server {
	return &Server{tripService: ts, Router: mux}
}