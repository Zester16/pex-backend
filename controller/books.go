package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"pex.oschmid.com/services"
)

// search books
func SearchBooks(ctx *fiber.Ctx) error {

	fmt.Println("into search books")
	query := ctx.Query("title")

	fmt.Println("Query", query)
	data, err := services.GetBooksFromOpenLibrary(query)

	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{"statusCode": 400, "error": err})
	}

	return ctx.JSON(&fiber.Map{"statusCode": 0, "data": data})
}
