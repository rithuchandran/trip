package service

import (
	"big-life-backend/trip/data"
	"big-life-backend/trip/domain"
)

type TripService struct {
	repository *data.TripRepository
}

type TripServiceInt interface {
	GetTrip(int) (domain.Trip, error)
	CreateTrip(domain.Trip)
	UpdateTrip(int, domain.Trip)
	DeleteTrip(int) error
}

func NewTripService(pr *data.TripRepository) *TripService {
	return &TripService{repository: pr}
}

func (ts TripService) GetTrip(id int) (domain.Trip, error) {
	return ts.repository.GetTrip(id)

}

func (ts TripService) CreateTrip(trip domain.Trip) {
	ts.repository.CreateTrip(trip)
}

func (ts TripService) UpdateTrip(id int, trip domain.Trip) {
	ts.repository.UpdateTrip(id, trip)
}

func (ts TripService) DeleteTrip(id int) error {
	return ts.repository.DeleteTrip(id)
}
