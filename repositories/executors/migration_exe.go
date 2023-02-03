package executors

import (
	"context"
	"log"

	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"gorm.io/gorm"
)

type migrationExecutors struct {
	*gorm.DB
}

func NewMigrationExecutor(DB *gorm.DB) repositories.MigrationRepo {
	return &migrationExecutors{DB: DB}
}

func (repo *migrationExecutors) InsertCountries(ctx context.Context, country []entities.Country) error {
	for i := 0; i < len(country); i++ {
		// log.Println(country[i].CallingCode)
		err := repo.DB.WithContext(ctx).Create(&country).Error
		if err != nil {
			log.Println(country[i].Id, '-', country[i].Name)
			exceptions.PanicLogging(err)
		}
	}
	return nil
}
