package main

import (
	"trip/data"
	"trip/routing"
	"trip/service"

	"github.com/gorilla/mux"
)

func main() {
	personRepository := data.NewPersonRepository()

	tripService := service.NewTripService(personRepository)

	router := mux.NewRouter()

	server := routing.NewServer(tripService, router)

	server.Start()
}
