package data

import (
	"big-life-backend/trip/domain"
)

//PersonRepository contains a slice of Trip
type PersonRepository struct {
	trips []domain.Trip
}

var sampleTrips = []domain.Trip{
	{Id: 1, Origin: "Bangalore", Destination: "Mysore",},
	{Id: 2, Origin: "Bangalore", Destination: "Chennai",},
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{trips: sampleTrips}
}
