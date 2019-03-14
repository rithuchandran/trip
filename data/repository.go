package data

import (
	"github.com/rithuchandran/trip/domain"
)

//TripRepository contains a slice of Trip
type TripRepository struct {
	trips []domain.Trip
}

var sampleTrips = []domain.Trip{
	{Id: 1, Origin: "Bangalore", Destination: "Mysore"},
	{Id: 2, Origin: "Bangalore", Destination: "Chennai"},
	{Id: 3, Origin: "Bangalore", Destination: "Goa"},
	{Id: 4, Origin: "Andaman", Destination: "Bangalore"},
	{Id: 5, Origin: "Pondichery", Destination: "Chennai"},
	{Id: 6, Origin: "Meghalaya", Destination: "Mysore"},
	{Id: 7, Origin: "Gangtok", Destination: "Chennai"},
	{Id: 8, Origin: "Bangalore", Destination: "Sikkim"},
	{Id: 9, Origin: "Bangalore", Destination: "Jaipur"},
	{Id: 10, Origin: "Bangalore", Destination: "Kochi"},
	{Id: 11, Origin: "Bangalore", Destination: "Hyderabad"},
	{Id: 12, Origin: "Kodaikanal", Destination: "Chennai"},
}

func NewTripRepository() *TripRepository {
	return &TripRepository{trips: sampleTrips}
}
