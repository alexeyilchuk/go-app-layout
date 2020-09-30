package models


import "time"

type PassengerID int64

type Passenger struct {
	ID         PassengerID
	Name       string
	createdAt  time.Time
	lastFlight time.Time
	deletedAt  *time.Time
	status     string
}

func (p Passenger) IsActive() bool {
	return p.deletedAt == nil && p.lastFlight.Before(time.Now().AddDate(-1, 0, 0))
}
