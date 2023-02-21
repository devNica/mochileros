package controllers

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/services"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	services.UserService
	configurations.Config
}

func NewUserController(service *services.UserService, config configurations.Config) *userController {
	return &userController{UserService: *service, Config: config}
}

func (controller userController) Route(app *fiber.App) {
	app.Post("/mochileros/v1/user/:userId/kyc", controller.RegisterKYC)
	app.Put("/mochileros/v1/user/:userId/status", controller.ChangeAccountStatus)
}

func (controller userController) RegisterKYC(c *fiber.Ctx) error {

	var request request.KYCRequestModel
	request.UserId = c.Params("userID")
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	controller.UserService.RegisterKYC(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "KYC registered successfully",
		Data:    "",
	})
}

func (controller userController) ChangeAccountStatus(c *fiber.Ctx) error {
	userId := c.Params("userID")

	response := controller.UserService.ChangeAccountStatus(c.Context(), userId)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}
