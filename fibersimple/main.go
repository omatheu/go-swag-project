package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/omatheu/go-swag-project/docs/fibersimple" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	"log"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	// Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// New Routes
	app.Get("/users", GetUsers)    // Get list of users
	app.Post("/users", CreateUser) // Create a new user

	// Start Server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

// GetUsers godoc
// @Summary Get list of users
// @Description Get a list of users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
	users := []map[string]interface{}{
		{"id": 1, "name": "John Doe"},
		{"id": 2, "name": "Jane Doe"},
	}
	return c.JSON(users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body map[string]interface{} true "User"
// @Success 201 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(map[string]interface{})
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
