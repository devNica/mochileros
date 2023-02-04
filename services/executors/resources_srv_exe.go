package executors

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
)

type resourcesServiceExecutor struct {
	repositories.CountryRepo
}

func NewResourcesServiceExecutor(repo *repositories.CountryRepo) services.ResourcesService {
	return &resourcesServiceExecutor{CountryRepo: *repo}
}

func (repo *resourcesServiceExecutor) GetAll(ctx context.Context) []response.CountryResponseModel {

	countries, err := repo.CountryRepo.FetchAll(ctx)
	exceptions.PanicLogging(err)

	// var countries []models.Country
	// for _, country := range res {

	// 	countries = append(countries, models.Country{
	// 		Id:          country.Id,
	// 		Name:        country.Name,
	// 		Capital:     country.Name,
	// 		Cca3:        country.Cca3,
	// 		CallingCode: country.CallingCode,
	// 		TimeZones:   country.TimeZones,
	// 		States:      country.States,
	// 		Latitude:    country.Latitude,
	// 		Longitude:   country.Longitude,
	// 		FlagPng:     country.FlagPng,
	// 		FlagSvg:     country.FlagSvg,
	// 		CurrCode:    country.CurrCode,
	// 		CurrName:    country.CurrName,
	// 		CurrSymbol:  country.CurrSymbol,
	// 	})
	// }

	return countries

}

func (repo *resourcesServiceExecutor) GetCountryByName(ctx context.Context, countryName string) response.CountryResponseModel {

	country, err := repo.CountryRepo.FetchCountryByName(ctx, countryName)
	exceptions.PanicLogging(err)

	return country

}
