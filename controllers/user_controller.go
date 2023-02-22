package controllers

import (
	"errors"

	"github.com/devNica/mochileros/commons/serializers"
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
}

func (controller userController) RegisterKYC(c *fiber.Ctx) error {

	var kycRequest request.KYCRequestModel
	kycRequest.UserId = c.Params("userID")
	err := c.BodyParser(&kycRequest)
	exceptions.PanicLogging(err)

	form, multipartError := c.MultipartForm()

	if multipartError != nil {
		exceptions.PanicLogging(multipartError)
	}

	files := form.File

	if len(files) == 0 {
		exceptions.PanicLogging(errors.New("Images not found"))
	}

	images, serializeError := serializers.SerializeUserAsset(files)
	if serializeError != nil {
		exceptions.PanicLogging(serializeError)
	}

	controller.UserService.RegisterKYC(c.Context(), kycRequest, images)

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "KYC registered successfully",
		Data:    "",
	})
}
