package exceptions

import (
	"encoding/json"

	"github.com/devNica/mochileros/models"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)

	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)

		return ctx.Status(fiber.StatusBadRequest).JSON(models.GeneralResponseModel{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})

	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(models.GeneralResponseModel{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusNotFound).JSON(models.GeneralResponseModel{
			Code:    401,
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(models.GeneralResponseModel{
		Code:    500,
		Message: "Internal Server Error",
		Data:    err.Error(),
	})
}
