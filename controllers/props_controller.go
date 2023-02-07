package controllers

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/services"
	"github.com/gofiber/fiber/v2"
)

type propsController struct {
	services.PropsService
	configurations.Config
}

func NewPropsController(srv *services.PropsService, config configurations.Config) *propsController {
	return &propsController{PropsService: *srv, Config: config}
}

func (controller propsController) Route(app *fiber.App) {
	app.Get("/mochileros/v1/props/country", controller.GetAllCountries)
	app.Get("/mochileros/v1/props/country/:name/", controller.GetCountryByName)
	app.Get("/mochileros/v1/props/migrate/country", controller.MigrateCountryInfo)
}

func (controller propsController) GetAllCountries(c *fiber.Ctx) error {

	response := controller.PropsService.GetAll(c.Context())
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    response,
	})
}

func (controller propsController) GetCountryByName(c *fiber.Ctx) error {

	response := controller.PropsService.GetCountryByName(c.Context(), c.Params("name"))
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    response,
	})
}

func (controller propsController) MigrateCountryInfo(c *fiber.Ctx) error {

	controller.PropsService.MigrateCountryInfo(c.Context())
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Countries has been migrated",
		Data:    "",
	})
}
