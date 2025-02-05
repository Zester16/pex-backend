package routes

import (
	"github.com/gofiber/fiber/v2"
	"pex.oschmid.com/controller"
)

func RoutesSplash(app *fiber.App) {
	api := app.Group("/v1/splash")

	api.Get("/", controller.GetAllSplash)

	api.Get("/:id", controller.GetSplash)
	api.Post("/:id", controller.AddSplash)
}
