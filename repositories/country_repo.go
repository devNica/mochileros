package repositories

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
)

type CountryRepo interface {
	InsertCountries(ctx context.Context, entity []entities.Country) error
	FetchAll(ctx context.Context) ([]response.CountryResponseModel, error)
	FetchCountryByName(ctx context.Context, countryName string) (response.CountryResponseModel, error)
}
