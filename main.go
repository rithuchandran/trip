package main

import (
	"bitbucket.org/big_life/big-life-backend/trip/data"
	"bitbucket.org/big_life/big-life-backend/trip/routing"
	"bitbucket.org/big_life/big-life-backend/trip/service"

	"github.com/gorilla/mux"
)

func main() {
	tripRepository := data.NewTripRepository()

	tripService := service.NewTripService(tripRepository)

	router := mux.NewRouter()

	server := routing.NewServer(tripService, router)

	server.Start()
}
