package app

import "time"

type FlightNumber string

type Flight struct {
	Number         FlightNumber
	DeparturePoint string
	ArrivalPoint   string
	Date           time.Time
	Status         string
}
