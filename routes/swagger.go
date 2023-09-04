package routes

import (
	_ "fiber-mongo-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

//	@title			Fiber MongoDB Api
//	@version		1.0
//	@description	This is a simple CRUD made with GO, Fiber Frameword and MongoDB.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Wesley Lima
//	@contact.url	https://github.com/wesleylimabh
//	@contact.email	email@email.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/

////	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func SwaggerRoute(app *fiber.App) {

	app.Get("/swagger/*", swagger.HandlerDefault)

}
