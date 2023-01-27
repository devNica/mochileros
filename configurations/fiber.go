package configurations

import (
	"github.com/devNica/mochileros/exceptions"
	"github.com/gofiber/fiber/v2"
)

func NewFiber() fiber.Config {
	return fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
	}
}
