package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pex.oschmid.com/model"
	"pex.oschmid.com/repository"
)

// *****NEWSPAPER ROUTES
// Adds newpaper
func AddNewspaper(c *fiber.Ctx) error {
	p := new(model.NewspaperModel)
	err := c.BodyParser(p)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"status": 1,

			"message": err,
		})

	}
	p.Id = uuid.New().String()
	err = repository.AddNewspaper(*p)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 400, "statusMessage": err.Error()})
	}
	return c.JSON(&fiber.Map{"statusCode": 0, "statusMessage": "success"})
}

//****READ NEWS

// Adds newspaper read date
func AddNewsRead(c *fiber.Ctx) error {
	p := new(model.NewspaperreadingModel)
	err := c.BodyParser(p)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"status": 1,

			"message": err,
		})

	}
	p.Id = uuid.New().String()
	err = repository.AddNewsRead(*p)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 400, "statusMessage": err.Error()})
	}
	return c.JSON(&fiber.Map{"statusCode": 0, "statusMessage": "success"})
}

func GetPaginatedNewsLettter(c *fiber.Ctx) error {
	return &fiber.Error{}
}
