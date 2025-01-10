package nosql

import (
	"fmt"

	"github.com/cifra-city/distributors-admin/internal/data/nosql/repositories"
)

type Repo struct {
	PlacesEmployees repositories.PlacesEmployees
}

func NewRepository(uri, dbName string) (*Repo, error) {
	placesEmployeesRepo, err := repositories.NewPlacesEmployees(uri, dbName, "places_employees")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize places_employees repository: %w", err)
	}

	return &Repo{
		PlacesEmployees: placesEmployeesRepo,
	}, nil
}
