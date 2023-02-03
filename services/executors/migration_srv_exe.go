package executors

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
)

func GetJson(url string, target interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// leer resultados de la peticion
	// content, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	return err
	// }

	// imprime por consola los resultados
	// log.Println(string(content))

	err = json.NewDecoder(resp.Body).Decode(target)

	if err != nil {
		return err
	}

	return nil
}

type migrationServiceExecutor struct {
	repositories.MigrationRepo
}

func NewMigrationServiceExecutor(repo *repositories.MigrationRepo) services.MigrationService {
	return &migrationServiceExecutor{MigrationRepo: *repo}
}

func (repo *migrationServiceExecutor) MigrateCountriesInfo(ctx context.Context) {

	country := []entities.Country{}

	url := "http://localhost:6700/migrate/country"
	err := GetJson(url, &country)

	exceptions.PanicLogging(err)

	repo.MigrationRepo.InsertCountries(ctx, country)

}
