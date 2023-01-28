package main

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/exceptions"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	//setup configurations
	config := configurations.New()
	configurations.DatabaseConnect(config)

	//setup fiber
	app := fiber.New(configurations.NewFiber())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	// start app
	err := app.Listen(config.Get("SERVER_PORT"))
	exceptions.PanicLogging(err)

}
