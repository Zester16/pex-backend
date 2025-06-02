package routes

import (
	"github.com/gofiber/fiber/v2"
	"pex.oschmid.com/controller"
)

func NewspaperRoutes(app *fiber.App) {

	api := app.Group("v1/newspaper")
	api.Post("/", controller.AddNewspaper)

}

func NewsReadRoutes(app *fiber.App) {
	api := app.Group("/v1/newsread")
	api.Group("/", controller.AddNewsRead)
}
