package executors

import (
	"context"
	"errors"

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

func (repo *countryRepoExecutor) FetchCountryByName(ctx context.Context, countryName string) (entities.Country, error) {

	var country entities.Country

	result := repo.DB.Where("name = ?", countryName).First(&country)

	if result.RowsAffected == 0 {
		return entities.Country{}, errors.New("Countries not found")
	}

	return country, nil
}
