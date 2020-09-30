package app

import "time"

type Registration struct {
	PassengerID  PassengerID
	FlightNumber FlightNumber
	CreatedAt    time.Time
	// e.t.c
}
