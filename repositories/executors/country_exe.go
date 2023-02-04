package executors

import (
	"context"
	"errors"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/repositories"
	"gorm.io/gorm"
)

type countryRepoExecutor struct {
	*gorm.DB
}

func NewCountryRepoExecutor(DB *gorm.DB) repositories.CountryRepo {
	return &countryRepoExecutor{DB: DB}
}

func (repo *countryRepoExecutor) FetchAll(ctx context.Context) ([]response.CountryResponseModel, error) {

	var listCountries []entities.Country

	result := repo.DB.Limit(20).Find(&listCountries)

	if result.RowsAffected == 0 {
		return []response.CountryResponseModel{}, errors.New("Countries not found")
	}

	var countries []response.CountryResponseModel
	for _, country := range listCountries {

		countries = append(countries, response.CountryResponseModel{
			Id:          country.Id,
			Name:        country.Name,
			Capital:     country.Name,
			Cca3:        country.Cca3,
			CallingCode: country.CallingCode,
			TimeZones:   country.TimeZones,
			States:      country.States,
			Latitude:    country.Latitude,
			Longitude:   country.Longitude,
			FlagPng:     country.FlagPng,
			FlagSvg:     country.FlagSvg,
			CurrCode:    country.CurrCode,
			CurrName:    country.CurrName,
			CurrSymbol:  country.CurrSymbol,
		})
	}

	return countries, nil
}

func (repo *countryRepoExecutor) FetchCountryByName(ctx context.Context, countryName string) (response.CountryResponseModel, error) {

	var foundCountry entities.Country

	result := repo.DB.Where("name = ?", countryName).First(&foundCountry)

	if result.RowsAffected == 0 {
		return response.CountryResponseModel{}, errors.New("Countries not found")
	}

	country := response.CountryResponseModel{
		Id:          foundCountry.Id,
		Name:        foundCountry.Name,
		Capital:     foundCountry.Capital,
		Cca3:        foundCountry.Cca3,
		CallingCode: foundCountry.CallingCode,
		TimeZones:   foundCountry.TimeZones,
		States:      foundCountry.States,
		Latitude:    foundCountry.Latitude,
		Longitude:   foundCountry.Longitude,
		FlagPng:     foundCountry.FlagPng,
		FlagSvg:     foundCountry.FlagSvg,
		CurrCode:    foundCountry.CurrCode,
		CurrName:    foundCountry.CurrName,
		CurrSymbol:  foundCountry.CurrSymbol,
	}

	return country, nil
}
