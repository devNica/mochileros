package main

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/controllers"
	"github.com/devNica/mochileros/exceptions"
	repository "github.com/devNica/mochileros/repositories/executors"
	service "github.com/devNica/mochileros/services/executors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	//setup configurations
	config := configurations.New()
	conn := configurations.DatabaseConnect(config)
	argon := configurations.NewArgonConfg()

	// repositories
	userRepository := repository.NewUserRepoExecutor(conn)
	countryRepository := repository.NewCountryRepoExecutor(conn)
	hotelRepository := repository.NewHotelRepoExecutor(conn)
	fileRepository := repository.NewFileRepositoryExecutor(conn)

	//services
	UserAccountService := service.NewUserSrvExecutor(&userRepository)
	AuthService := service.NewAuthSrvExecutor(&userRepository, &argon)
	ResourcesService := service.NewResourcesServiceExecutor(&countryRepository, &fileRepository)
	HotelService := service.NewHotelServiceExecutor(&hotelRepository, &fileRepository)

	//controllers
	authController := controllers.NewAuthController(&AuthService, config)
	userController := controllers.NewUserController(&UserAccountService, config)
	propsController := controllers.NewPropsController(&ResourcesService, config)
	hotelController := controllers.NewOwnerController(&HotelService, config)

	//setup fiber
	app := fiber.New(configurations.NewFiber())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	// routing
	authController.Route(app)
	userController.Route(app)
	propsController.Route(app)
	hotelController.Route(app)

	// start app
	err := app.Listen(config.Get("SERVER_PORT"))
	exceptions.PanicLogging(err)

}
