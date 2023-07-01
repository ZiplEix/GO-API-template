package main

import (
	"fmt"
	"os"

	"github.com/ZiplEix/API_template/config"
	"github.com/ZiplEix/API_template/internal/storage"
	"github.com/ZiplEix/API_template/internal/todo"
	"github.com/ZiplEix/API_template/internal/user"
	"github.com/ZiplEix/API_template/pkg/shutdown"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"

	_ "github.com/ZiplEix/API_template/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title API Template
// @version 2.0
// @description An example template of a Golang backend API using Fiber and Postgres.
// @contact.name ZiplEix
// host localhost:3000
// @BasePath /
func main() {
	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error: ", err)
		exitCode = 1
		return
	}

	// run the server
	cleanup, err := run(env)

	// run the cleanup after the server is terminated
	defer cleanup()
	if err != nil {
		fmt.Println("error: ", err)
		exitCode = 1
		return
	}

	// ensure the server is shutdown gracefully and run app
	shutdown.Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	// start the server
	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	// return a function to close the server and database
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) (*fiber.App, func(), error) {
	// init database
	db, err := storage.BootstrapPostgres(env.DB_USER, env.DB_PASSWORD, env.DB_NAME)
	if err != nil {
		return nil, nil, err
	}

	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// add routes
	router(app, db)

	// return the app and a cleanup function
	return app, func() {
		storage.ClosePostgres(db)
	}, nil
}

func router(app *fiber.App, db *gorm.DB) {
	// add ping check
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// add docs
	app.Get("/swagger/*", swagger.HandlerDefault)

	// user routes
	userStore := user.NewUserStorage(db)
	userController := user.NewUserController(userStore)
	user.AddUserRoutes(app, userController)

	// todo routes
	todoStore := todo.NewTodoStorage(db)
	todoController := todo.NewTodoController(todoStore)
	todo.AddTodoRoutes(app, todoController, userController)

}
