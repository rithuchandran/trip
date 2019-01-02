package main

import (
	"big-life-backend/trip/data"
	"big-life-backend/trip/routing"
	"big-life-backend/trip/service"

	"github.com/gorilla/mux"
)

func main() {
	tripRepository := data.NewTripRepository()

	tripService := service.NewTripService(tripRepository)

	router := mux.NewRouter()

	server := routing.NewServer(tripService, router)

	server.Start()
}
