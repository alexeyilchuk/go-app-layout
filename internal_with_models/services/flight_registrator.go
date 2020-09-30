package services

import (
	"time"

	"github.com/alexeyilchuk/go-app-layout/internal_with_models/models"
)

type airlineRegistrationService interface {
	Register(flight models.Flight, passengerName string) (id string, err error)
}

type registrationStorage interface {
	Create(registration models.Registration) error
}

type FlightRegister struct {
	airlineService      airlineRegistrationService
	registrationStorage registrationStorage
}

func (fr FlightRegister) RegisterOnFlight(flight models.Flight, passenger models.Passenger) error {
	_, err := fr.airlineService.Register(flight, passenger.Name)
	if err != nil {
		return err
	}

	return fr.registrationStorage.Create(models.Registration{
		PassengerID:  passenger.ID,
		FlightNumber: flight.Number,
		CreatedAt:    time.Now(),
	})
}
