package app

import "time"

type airlineRegistrationService interface {
	Register(flight Flight, passengerName string) (id string, err error)
}

type registrationStorage interface {
	Create(registration Registration) error
}

type FlightRegister struct {
	airlineService      airlineRegistrationService
	registrationStorage registrationStorage
}

func (fr FlightRegister) RegisterOnFlight(flight Flight, passenger Passenger) error {
	_, err := fr.airlineService.Register(flight, passenger.Name)
	if err != nil {
		return err
	}

	return fr.registrationStorage.Create(Registration{
		PassengerID:  passenger.ID,
		FlightNumber: flight.Number,
		CreatedAt:    time.Now(),
	})
}
