package services

import (
	"testing"

	"github.com/alexeyilchuk/go-app-layout/internal_with_models/models"
)

func TestFlightRegister_RegisterOnFlight(t *testing.T) {
	type fields struct {
		airlineService      airlineRegistrationService
		registrationStorage registrationStorage
	}
	type args struct {
		flight    models.Flight
		passenger models.Passenger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := FlightRegister{
				airlineService:      tt.fields.airlineService,
				registrationStorage: tt.fields.registrationStorage,
			}
			if err := fr.RegisterOnFlight(tt.args.flight, tt.args.passenger); (err != nil) != tt.wantErr {
				t.Errorf("RegisterOnFlight() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
