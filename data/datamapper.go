package data

import (
	"errors"
	"github.com/rithuchandran/trip/domain"
)

func (pr *TripRepository) GetTrip(id int) (domain.Trip, error) {
	tr, _, err := findTrip(pr, id)
	return tr, err
}

func (pr *TripRepository) CreateTrip(trip domain.Trip) {
	pr.trips = append(pr.trips, trip)
}
func (pr *TripRepository) UpdateTrip(id int, trip domain.Trip) {
	for i, tr := range pr.trips {
		if id == tr.Id {
			pr.trips[i] = trip
			pr.trips[i].Id = id
			return
		}
	}
}

func (pr *TripRepository) DeleteTrip(id int) error {
	_, i, err := findTrip(pr, id)
	if err != nil {
		return err
	}
	pr.trips = append((pr.trips)[:i], (pr.trips)[i+1:]...)
	return nil
}

func findTrip(pr *TripRepository, id int) (domain.Trip, int, error) {
	for i, trip := range pr.trips {
		if id == trip.Id {
			return trip, i, nil
		}
	}
	return domain.Trip{}, -1, errors.New("trip does not exist")
}
