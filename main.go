package main

import (
	"bitbucket.org/big_life/big-life-backend/trip/data"
	"bitbucket.org/big_life/big-life-backend/trip/routing"
	"bitbucket.org/big_life/big-life-backend/trip/service"

	"github.com/gorilla/mux"
)

func main() {
	personRepository := data.NewPersonRepository()

	tripService := service.NewTripService(personRepository)

	router := mux.NewRouter()

	server := routing.NewServer(tripService, router)

	server.Start()
}
