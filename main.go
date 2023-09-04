package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	// run database
	configs.ConnectDB()

	// routes
	routes.UserRoute(app)
	routes.SwaggerRoute(app)

	app.Listen(":3000")

}
