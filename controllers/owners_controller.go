package controllers

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/services"
	"github.com/gofiber/fiber/v2"
)

type hotelController struct {
	services.HotelService
	configurations.Config
}

func NewOwnerController(srv *services.HotelService, config configurations.Config) *hotelController {
	return &hotelController{HotelService: *srv, Config: config}
}

func (controller hotelController) Route(app *fiber.App) {
	app.Post("/mochileros/v1/owner/hotel", controller.RegisterHotel)
	app.Get("/mochileros/v1/owner/:ownerId/hotel", controller.ListOwnerHotels)
}

func (controller hotelController) RegisterHotel(c *fiber.Ctx) error {

	var hotel request.HotelRequestModel

	err := c.BodyParser(&hotel)
	exceptions.PanicLogging(err)

	controller.HotelService.RegisterHotel(c.Context(), hotel)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    "",
	})
}

func (controller hotelController) ListOwnerHotels(c *fiber.Ctx) error {

	response := controller.HotelService.GetListOwnerHotels(c.Context(), c.Params("ownerId"))
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    response,
	})
}
