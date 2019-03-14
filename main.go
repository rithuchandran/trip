package main

import (
	"github.com/rithuchandran/trip/data"
	"github.com/rithuchandran/trip/logger"
	"github.com/rithuchandran/trip/routing"
	"github.com/rithuchandran/trip/service"

	"github.com/gorilla/mux"
)

func main() {

	tripRepository := data.NewTripRepository()

	tripService := service.NewTripService(tripRepository)

	router := mux.NewRouter()

	server := routing.NewServer(tripService, router, logger.NewLogger())

	server.Start()
}
