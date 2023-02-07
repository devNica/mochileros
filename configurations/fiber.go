package configurations

import (
	"github.com/devNica/mochileros/exceptions"
	"github.com/gofiber/fiber/v2"
)

func NewFiber() fiber.Config {
	return fiber.Config{
		BodyLimit:    5 * 1024 * 1024, //Limit upload file 5mb
		ErrorHandler: exceptions.ErrorHandler,
	}
}
