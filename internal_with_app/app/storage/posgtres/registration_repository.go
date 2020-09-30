package posgtres

import "github.com/alexeyilchuk/go-app-layout/internal_with_app/app"

type RegistrationPGRepository struct {
	// db connect
}

func (r RegistrationPGRepository) Create(f app.Flight) error {
	return nil
}
