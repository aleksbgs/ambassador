package main

import (
	"github.com/aleksbgs/ambassador/src/database"
	"github.com/aleksbgs/ambassador/src/routes"
	"github.com/aleksbgs/ambassador/src/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()
	database.SetupCacheChannel()
	services.Setup()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
