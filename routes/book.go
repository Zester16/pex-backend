package routes

import (
	"github.com/gofiber/fiber/v2"
	"pex.oschmid.com/controller"
)

// Handles all requests related to books
// aka searching, adding, adding comments, mark as now reading, completed
func BookRoute(app *fiber.App) {

	api := app.Group("/v1/books")
	api.Get("/", controller.SearchBooks)

}
