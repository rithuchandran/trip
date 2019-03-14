package routing

import (
	"cloud.google.com/go/logging"
	"github.com/rithuchandran/trip/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	tripService service.TripServiceInt
	Router      *mux.Router
	logger      *logging.Logger
}

func (s *Server) Start() {
	s.Routes()
	_ = http.ListenAndServe(":8080", s.Router)
}

func NewServer(ts service.TripServiceInt, mux *mux.Router, lg *logging.Logger) *Server {
	return &Server{tripService: ts, Router: mux}
}
