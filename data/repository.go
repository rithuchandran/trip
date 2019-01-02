package data

import (
	"bitbucket.org/big_life/big-life-backend/trip/domain"
)

//TripRepository contains a slice of Trip
type TripRepository struct {
	trips []domain.Trip
}

var sampleTrips = []domain.Trip{
	{Id: 1, Origin: "Bangalore", Destination: "Mysore"},
	{Id: 2, Origin: "Bangalore", Destination: "Chennai"},
}

func NewTripRepository() *TripRepository {
	return &TripRepository{trips: sampleTrips}
}
