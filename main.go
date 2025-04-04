package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/joho/godotenv"
	"pex.oschmid.com/controller"
	"pex.oschmid.com/database"
	"pex.oschmid.com/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	deployEnv := os.Getenv("DEPLOY_ENV")
	app := fiber.New()
	if deployEnv == "dev" {
		app.Use(cors.New(
			cors.Config{
				AllowOrigins:     "http://localhost:3000",
				AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
				AllowMethods:     "GET,PUT,POST,DELETE,PATCH,OPTIONS",
				AllowCredentials: true,
			}))
	} else {
		app.Use(cors.New(
			cors.Config{
				AllowOrigins:     "http://localhost:3000",
				AllowHeaders:     "Origin, Content-Type, Accept",
				AllowCredentials: true,
			}))

	}
	sessKey := os.Getenv("COOKIE_KEY")
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: sessKey,
	}))
	err = database.ConnectDB()

	if err != nil {
		log.Fatal("Database connection failed", err)
	}
	app.Get("/test", controller.MiddlewareCheckUser, func(c *fiber.Ctx) error {
		fmt.Println(string(c.Context().UserAgent()))
		return c.JSON(&fiber.Map{
			"test":       "hi",
			"user-agent": string(c.Context().UserAgent()),
		})

	})

	//external app routings
	//routing for splash /v1/splash/
	routes.RoutesSplash(app)
	// routing for user /v1/user
	routes.RoutesUser(app)

	log.Fatal(app.Listen(":4000"))
}
