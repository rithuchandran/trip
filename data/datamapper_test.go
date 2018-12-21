package data_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trip/data"
	"trip/domain"
)

func TestPersonRepository_GetTrip(t *testing.T) {
	tr, _ := data.NewPersonRepository().GetTrip(1)
	assert.Equal(t, "Bangalore", tr.Origin, nil)
	assert.Equal(t, "Mysore", tr.Destination, nil)

}

func TestPersonRepository_CreateTrip(t *testing.T) {
	trip := domain.Trip{Id: 3, Destination: "SomeWhere", Origin: "Here"}
	data.NewPersonRepository().CreateTrip(trip)
	tr, _ := data.NewPersonRepository().GetTrip(3)
	assert.Equal(t, trip, tr, nil)
}

func TestPersonRepository_UpdateTrip(t *testing.T) {
	trip := domain.Trip{Destination: "abc", Origin: "def"}
	data.NewPersonRepository().UpdateTrip(2, trip)
	tr, _ := data.NewPersonRepository().GetTrip(2)
	assert.Equal(t, trip.Origin, tr.Origin, nil)
}

//noinspection GoUnhandledErrorResult
func TestPersonRepository_DeleteTrip(t *testing.T) {
	p := data.NewPersonRepository()
	_, e := p.GetTrip(1)
	assert.Equal(t, nil, e, nil)
	data.NewPersonRepository().DeleteTrip(1)
	_, e2 := p.GetTrip(1)
	assert.NotEqual(t, nil, e2, nil)
}