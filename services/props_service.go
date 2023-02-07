package services

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
)

type PropsService interface {
	GetAll(ctx context.Context) []response.CountryResponseModel
	GetCountryByName(ctx context.Context, countryName string) response.CountryResponseModel
	MigrateCountryInfo(ctx context.Context)
}
