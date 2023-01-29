package controllers

import (
	"fmt"

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
	app.Post("/mochileros/v1/account/user", controller.GetUserByEmail)
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

func (controller userAccountController) GetUserByEmail(c *fiber.Ctx) error {

	type fetchUser struct {
		Email string `json:"email" validate:"required"`
	}

	request := fetchUser{}

	fmt.Println("request", c.BodyParser(&request))

	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	response := controller.UserAccountService.GetUserByEmail(c.Context(), request.Email)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "User has been registered successfull",
		Data:    response,
	})

}
