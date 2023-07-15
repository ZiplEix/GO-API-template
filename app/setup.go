package app

import (
	"os"

	"github.com/ZiplEix/API_template/config"
	"github.com/ZiplEix/API_template/database"
	"github.com/ZiplEix/API_template/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() error {
	// Load ENV
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.ConnectDB()
	if err != nil {
		return err
	}

	// defer database close
	defer database.CloseDB()

	// create app
	app := fiber.New()

	// Set log for fiber
	app.Use(cors.New())
	app.Use(logger.New())

	// recover from panic
	app.Use(recover.New())

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// start app
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
