package data_test

import (
	"bitbucket.org/big_life/big-life-backend/trip/data"
	"bitbucket.org/big_life/big-life-backend/trip/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonRepository_GetTrip(t *testing.T) {
	tr, _ := data.NewTripRepository().GetTrip(1)
	assert.Equal(t, "Bangalore", tr.Origin, nil)
	assert.Equal(t, "Mysore", tr.Destination, nil)

}

func TestPersonRepository_CreateTrip(t *testing.T) {
	trip := domain.Trip{Id: 3, Destination: "SomeWhere", Origin: "Here"}
	personRepository := data.NewTripRepository()
	personRepository.CreateTrip(trip)
	tr, _ := personRepository.GetTrip(3)
	assert.Equal(t, trip, tr, nil)
}

func TestPersonRepository_UpdateTrip(t *testing.T) {
	trip := domain.Trip{Id:2,Destination: "abc", Origin: "def"}
	data.NewTripRepository().UpdateTrip(2, trip)
	tr, _ := data.NewTripRepository().GetTrip(2)

	assert.Equal(t, trip, tr, nil)
}

func TestPersonRepository_DeleteTrip(t *testing.T) {
	p := data.NewTripRepository()
	_, e := p.GetTrip(1)
	assert.Equal(t, nil, e, nil)
	data.NewTripRepository().DeleteTrip(1)
	_, e2 := p.GetTrip(1)
	assert.NotEqual(t, nil, e2, nil)
}
