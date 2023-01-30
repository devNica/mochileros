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
	app.Post("/mochileros/v1/account/user/:userId/kyc", controller.RegisterKYC)
	app.Get("/mochileros/v1/account/user/:userId", controller.GetCompleteUserInfo)
	app.Put("/mochileros/v1/account/user/:userId/status", controller.ChangeAccountStatus)
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

func (controller userAccountController) RegisterKYC(c *fiber.Ctx) error {

	var request models.KYCRequestModel
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

func (controller userAccountController) GetCompleteUserInfo(c *fiber.Ctx) error {

	userId := c.Params("userID")

	response := controller.UserAccountService.GetCompleteUserInfo(c.Context(), userId)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    200,
		Message: "Success",
		Data:    response,
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
