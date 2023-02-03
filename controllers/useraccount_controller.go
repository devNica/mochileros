package controllers

import (
	"log"

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
	app.Get("/mochileros/v1/account/user/:email", controller.GetUserByEmail)
	app.Post("/mochileros/v1/account/user/:userId/kyc", controller.RegisterKYC)
	app.Get("/mochileros/v1/account/user/:userId/info", controller.GetCompleteUserInfo)
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

	log.Println("es que entro aca")

	response := controller.UserAccountService.GetUserByEmail(c.Context(), c.Params("email"))

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
	log.Println("userId: ", userId)

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
