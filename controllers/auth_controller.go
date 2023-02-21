package controllers

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/services"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	services.AuthService
	configurations.Config
}

func NewAuthController(service *services.AuthService, config configurations.Config) *authController {
	return &authController{AuthService: *service, Config: config}
}

func (controller authController) Route(app *fiber.App) {
	app.Post("/mochileros/v1/auth/customer", controller.CustomerRegistration)
	app.Post("/mochileros/v1/auth/user/:userId/kyc", controller.RegisterKYC)
	app.Put("/mochileros/v1/auth/user/:userId/status", controller.ChangeAccountStatus)
	app.Post("/mochileros/v1/auth/login", controller.UserLogin)
}

func (controller authController) CustomerRegistration(c *fiber.Ctx) error {
	var request request.UserAccounRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	controller.AuthService.CustomerRegister(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Customer has been registered successfull",
		Data:    "",
	})
}

func (controller authController) RegisterKYC(c *fiber.Ctx) error {

	var request request.KYCRequestModel
	request.UserId = c.Params("userID")
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	controller.AuthService.RegisterKYC(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "KYC registered successfully",
		Data:    "",
	})
}

func (controller authController) ChangeAccountStatus(c *fiber.Ctx) error {
	userId := c.Params("userID")

	response := controller.AuthService.ChangeAccountStatus(c.Context(), userId)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller authController) UserLogin(c *fiber.Ctx) error {
	var request request.UserAccounRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	user := controller.AuthService.UserLogin(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    200,
		Message: "Success",
		Data:    user,
	})
}
