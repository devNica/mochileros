package controllers

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/services"
	"github.com/gofiber/fiber/v2"
)

type userAccountController struct {
	services.UserAccountService
	configurations.Config
}

func NewUserAccountController(service *services.UserAccountService, config configurations.Config) *userAccountController {
	return &userAccountController{UserAccountService: *service, Config: config}
}

func (controller userAccountController) Route(app *fiber.App) {
	app.Post("/mochileros/v1/account/customer", controller.Register)
	app.Post("/mochileros/v1/account/user/:userId/kyc", controller.RegisterKYC)
	app.Put("/mochileros/v1/account/user/:userId/status", controller.ChangeAccountStatus)
	app.Post("/mochileros/v1/account/login", controller.UserLogin)
}

func (controller userAccountController) Register(c *fiber.Ctx) error {
	var request request.UserAccounRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	controller.UserAccountService.UserAccountRegister(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "User has been registered successfull",
		Data:    "",
	})
}

func (controller userAccountController) RegisterKYC(c *fiber.Ctx) error {

	var request request.KYCRequestModel
	request.UserId = c.Params("userID")
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	controller.UserAccountService.RegisterKYC(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "KYC registered successfully",
		Data:    "",
	})
}

func (controller userAccountController) ChangeAccountStatus(c *fiber.Ctx) error {
	userId := c.Params("userID")

	response := controller.UserAccountService.ChangeAccountStatus(c.Context(), userId)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller userAccountController) UserLogin(c *fiber.Ctx) error {
	var request request.UserAccounRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	user := controller.UserAccountService.UserLogin(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    200,
		Message: "Success",
		Data:    user,
	})
}
