package main

import (
	"fiber-mongo-api/configs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// run database
	configs.ConnectDB()

	app.Listen(":3000")

}
