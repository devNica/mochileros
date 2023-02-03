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
	userAccountRepository := repository.NewUserAccountExecutor(conn)
	migrationRepository := repository.NewMigrationExecutor(conn)
	countryRepository := repository.NewCountryRepoExecutor(conn)
	hotelRepository := repository.NewHotelRepoExecutor(conn)

	//services
	UserAccountService := service.NewUserAccountSrvExecutor(&userAccountRepository, &argon)
	MigrationService := service.NewMigrationServiceExecutor(&migrationRepository)
	ResourcesService := service.NewResourcesServiceExecutor(&countryRepository)
	HotelService := service.NewHotelServiceExecutor(&hotelRepository)

	//controllers
	userAccountController := controllers.NewUserAccountController(&UserAccountService, config)
	migrationController := controllers.NewMigrationController(&MigrationService, config)
	propsController := controllers.NewPropsController(&ResourcesService, config)
	hotelController := controllers.NewhotelController(&HotelService, config)

	//setup fiber
	app := fiber.New(configurations.NewFiber())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	// routing
	userAccountController.Route(app)
	migrationController.Route(app)
	propsController.Route(app)
	hotelController.Route(app)

	// start app
	err := app.Listen(config.Get("SERVER_PORT"))
	exceptions.PanicLogging(err)

}
