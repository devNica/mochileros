package controllers

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/services"
	"github.com/gofiber/fiber/v2"
)

type hotelController struct {
	services.HotelService
	configurations.Config
}

func NewhotelController(srv *services.HotelService, config configurations.Config) *hotelController {
	return &hotelController{HotelService: *srv, Config: config}
}

func (controller hotelController) Route(app *fiber.App) {
	app.Post("/mochileros/v1/user/:userId/hotel", controller.RegisterHotel)
	app.Get("/mochileros/v1/user/:userId/hotel", controller.GetAllByOwnerId)
}

func (controller hotelController) RegisterHotel(c *fiber.Ctx) error {

	var hotel models.HotelRequestModel

	err := c.BodyParser(&hotel)
	exceptions.PanicLogging(err)

	hotel.OwnerId = c.Params("userId")

	controller.HotelService.RegisterHotel(c.Context(), hotel)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    "",
	})
}

func (controller hotelController) GetAllByOwnerId(c *fiber.Ctx) error {

	response := controller.HotelService.GetAllByOwnerId(c.Context(), c.Params("userId"))
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    response,
	})
}
