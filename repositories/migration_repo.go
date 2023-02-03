package repositories

import (
	"context"

	"github.com/devNica/mochileros/entities"
)

type MigrationRepo interface {
	InsertCountries(ctx context.Context, entity []entities.Country) error
}
