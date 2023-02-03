package services

import "context"

type MigrationService interface {
	MigrateCountriesInfo(ctx context.Context)
}
