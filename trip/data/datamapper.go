package data

import (
	"errors"
	"big-life-backend/trip/domain"
)

func (pr PersonRepository) GetTrip(id int) (domain.Trip, error) {
	tr, _, err := findTrip(pr, id)
	return tr, err
}

func (pr PersonRepository) CreateTrip(trip domain.Trip) {
	*pr.trips = append(*pr.trips, trip)
}
func (pr PersonRepository) UpdateTrip(id int, trip domain.Trip) {
	for i, tr := range *pr.trips {
		if id == tr.Id {
			(*pr.trips)[i] = trip
			(*pr.trips)[i].Id = id
			return
		}
	}
}

func (pr PersonRepository) DeleteTrip(id int) error {
	_, i, err := findTrip(pr, id)
	if err != nil {
		return err
	}
	*pr.trips = append((*pr.trips)[:i], (*pr.trips)[i+1:]...)
	return nil
}

func findTrip(pr PersonRepository, id int) (domain.Trip, int, error) {
	for i, trip := range *pr.trips {
		if id == trip.Id {
			return trip, i, nil
		}
	}
	return domain.Trip{}, -1, errors.New("trip does not exist")
}
