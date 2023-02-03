package repositories

import (
	"context"

	"github.com/devNica/mochileros/entities"
)

type CountryRepo interface {
	FetchAll(ctx context.Context) ([]entities.Country, error)
	FetchCountryByName(ctx context.Context, countryName string) (entities.Country, error)
}
