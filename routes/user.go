package routes

import (
	"github.com/gofiber/fiber/v2"
	"pex.oschmid.com/controller"
)

func RoutesUser(app *fiber.App) {
	api := app.Group("/v1/user")
	api.Post("/login", controller.Login)
	api.Post("/refresh", controller.RegenerateToken)
	api.Post("/logout", controller.Logout)
	api.Get("/device", controller.MiddlewareCheckUser, controller.GetAllSessions)
}
