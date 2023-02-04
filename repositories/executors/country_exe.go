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

func (repo *countryRepoExecutor) FetchAll(ctx context.Context) ([]entities.Country, error) {

	var countries []entities.Country

	result := repo.DB.Find(&countries)

	if result.RowsAffected == 0 {
		return []entities.Country{}, errors.New("Countries not found")
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
