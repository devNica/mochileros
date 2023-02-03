package controllers

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/services"
	"github.com/gofiber/fiber/v2"
)

type migrationController struct {
	services.MigrationService
	configurations.Config
}

func NewMigrationController(srv *services.MigrationService, config configurations.Config) *migrationController {
	return &migrationController{MigrationService: *srv, Config: config}
}

func (controller migrationController) Route(app *fiber.App) {
	app.Get("/mochileros/v1/migration/countries", controller.MigrateCountries)
}

func (controller migrationController) MigrateCountries(c *fiber.Ctx) error {

	controller.MigrationService.MigrateCountriesInfo(c.Context())
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Countries has been migrated",
		Data:    "",
	})
}
