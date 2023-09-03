package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:userId", controllers.GetAUser)
	app.Put("/users/:userId", controllers.EditAUser)
	app.Delete("/users/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
}
