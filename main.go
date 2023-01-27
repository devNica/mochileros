package main

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	//setup configurations
	config := configurations.New()

	//setup fiber
	app := fiber.New(configurations.NewFiber())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	// start app
	err := app.Listen(config.Get("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

}
