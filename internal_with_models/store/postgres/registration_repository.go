package posgtres

import "github.com/alexeyilchuk/go-app-layout/internal_with_models/models"

type RegistrationPGRepository struct {
	// db connect
}

func (r RegistrationPGRepository) Create(f models.Flight) error {
	return nil
}
