package controllers

import (
	"github.com/devNica/mochileros/configurations"
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
}

func (controller userAccountController) Register(c *fiber.Ctx) error {
	var request models.UserAccounRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	controller.UserAccountService.UserAccountRegister(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "User has been registered successfull",
		Data:    "",
	})
}
