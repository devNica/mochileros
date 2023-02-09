package executors

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/devNica/mochileros/commons/encryptors/cipher"
	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
	"github.com/google/uuid"
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

type propsServiceExecutor struct {
	repositories.CountryRepo
	repositories.FileRepo
}

func NewResourcesServiceExecutor(Country *repositories.CountryRepo, File *repositories.FileRepo) services.PropsService {
	return &propsServiceExecutor{CountryRepo: *Country, FileRepo: *File}
}

func (repo *propsServiceExecutor) MigrateCountryInfo(ctx context.Context) {

	country := []entities.Country{}

	url := "http://localhost:6700/migrate/country"
	err := GetJson(url, &country)

	exceptions.PanicLogging(err)

	repo.CountryRepo.InsertCountries(ctx, country)

}

func (repo *propsServiceExecutor) GetAll(ctx context.Context) []response.CountryResponseModel {

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

func (repo *propsServiceExecutor) GetCountryByName(ctx context.Context, countryName string) response.CountryResponseModel {

	country, err := repo.CountryRepo.FetchCountryByName(ctx, countryName)
	exceptions.PanicLogging(err)

	return country

}

func (repo *propsServiceExecutor) DownloadHotelAsset(ctx context.Context, filekey string) response.FileResponseModel {
	str, err := cipher.Decrypt(filekey)
	exceptions.PanicLogging(err)

	s := strings.Split(str, ",")

	filename, err := uuid.Parse(s[0])

	file := request.FileDownloadRequestModel{
		Filename: filename,
		Filetype: s[1],
	}

	asset, err := repo.FileRepo.FetchHotelAsset(ctx, entities.File{
		Filename: file.Filename,
		Filetype: file.Filetype,
	})

	return response.FileResponseModel{
		Filetype: asset.Filetype,
		Binary:   asset.Binary,
	}
}
