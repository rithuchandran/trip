package routing

import (
	"bitbucket.org/big_life/big-life-backend/trip/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//noinspection GoUnhandledErrorResult
func (s Server) GetTripHandler(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id, _ := strconv.Atoi(parameters["id"])
	t, e := s.tripService.GetTrip(id)
	if e != nil {
		fmt.Fprint(writer, e)
		return
	}
	json.NewEncoder(writer).Encode(t)
}

//noinspection GoUnhandledErrorResult
func (s Server) CreateTripHandler(writer http.ResponseWriter, request *http.Request) {
	var trip domain.Trip
	err := json.NewDecoder(request.Body).Decode(&trip)
	trip.Id, _ = strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		fmt.Fprint(writer, "invalid input!")
		return
	}
	s.tripService.CreateTrip(trip)
	fmt.Fprint(writer, "Your trip has been created")
}

//noinspection GoUnhandledErrorResult
func (s Server) UpdateTripHandler(writer http.ResponseWriter, request *http.Request) {
	var trip domain.Trip
	err := json.NewDecoder(request.Body).Decode(&trip)
	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		fmt.Fprint(writer, "invalid input!")
		return
	}
	s.tripService.UpdateTrip(id, trip)
	fmt.Fprint(writer, "Your trip has been updated")
}

//noinspection GoUnhandledErrorResult
func (s Server) DeleteHandler(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id, _ := strconv.Atoi(parameters["id"])
	e := s.tripService.DeleteTrip(id)
	if e != nil {
		fmt.Fprint(writer, e)
		return
	}
	fmt.Fprint(writer, "Your trip has been deleted")
}
