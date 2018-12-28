package main

import (
	"big-life-backend/trip/data"
	"big-life-backend/trip/routing"
	"big-life-backend/trip/service"

	"github.com/gorilla/mux"
)

func main() {
	personRepository := data.NewPersonRepository()

	tripService := service.NewTripService(personRepository)

	router := mux.NewRouter()

	server := routing.NewServer(tripService, router)

	server.Start()
}
