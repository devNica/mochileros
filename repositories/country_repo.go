package repositories

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
)

type CountryRepo interface {
	FetchAll(ctx context.Context) ([]response.CountryResponseModel, error)
	FetchCountryByName(ctx context.Context, countryName string) (response.CountryResponseModel, error)
}
