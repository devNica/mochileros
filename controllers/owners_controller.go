package controllers

import (
	"bytes"
	"io"

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

func NewOwnerController(hotelSrv *services.HotelService, config configurations.Config) *hotelController {
	return &hotelController{HotelService: *hotelSrv, Config: config}
}

func (controller hotelController) Route(app *fiber.App) {
	app.Post("/mochileros/v1/owner/hotel", controller.RegisterHotel)
	app.Get("/mochileros/v1/owner/:ownerId/hotel", controller.ListOwnerHotels)
}

func (controller hotelController) RegisterHotel(c *fiber.Ctx) error {

	var hotel request.HotelRequestModel
	var fileReq request.FileRequestModel

	file, paramsFileErr := c.FormFile("hotelAsset")
	exceptions.PanicLogging(paramsFileErr)

	buffer, bufferErr := file.Open()
	exceptions.PanicLogging(bufferErr)
	defer buffer.Close()

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, buffer)

	fileReq.Buffer = buf.Bytes()
	fileReq.Filetype = file.Header.Get("Content-Type")
	fileReq.Filesize = int(file.Size)

	parserErr := c.BodyParser(&hotel)
	exceptions.PanicLogging(parserErr)

	controller.HotelService.RegisterHotel(c.Context(), hotel, fileReq)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    "",
	})
}

func (controller hotelController) ListOwnerHotels(c *fiber.Ctx) error {

	response := controller.HotelService.GetListOwnerHotels(c.Context(), c.BaseURL(), c.Params("ownerId"))
	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponseModel{
		Code:    201,
		Message: "Sucessfull Requets",
		Data:    response,
	})
}
