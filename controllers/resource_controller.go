package controllers

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/services"
	"github.com/gofiber/fiber/v2"
)

type propsController struct {
	services.ResourcesService
	configurations.Config
}

func NewPropsController(srv *services.ResourcesService, config configurations.Config) *propsController {
	return &propsController{ResourcesService: *srv, Config: config}
}

func (controller propsController) Route(app *fiber.App) {
	app.Get("/mochileros/v1/props/country", controller.GetAllCountries)
}

func (controller propsController) GetAllCountries(c *fiber.Ctx) error {

	response := controller.ResourcesService.GetAll(c.Context())
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    response,
	})
}
