package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"pex.oschmid.com/database"
	"pex.oschmid.com/model"
	"pex.oschmid.com/repository"
)

func GetSplash(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	fmt.Println("GetSplash params", id)
	if err != nil {
		fmt.Println("splash-controller-error:", err)
		return c.JSON(&fiber.Map{"statusCode": 400, "message": "missing path param id"})
	}

	resp, err := repository.GetIndividualSplash(id)

	if err != nil {
		fmt.Println("splash-controller-error:", err)
		return c.Status(400).JSON(&fiber.Map{"statusCode": 400, "error": err})
	}
	return c.JSON(&fiber.Map{"statusCode": 200, "data": resp})
}

func AddSplash(c *fiber.Ctx) error {
	p := new(model.Splash)
	err := c.BodyParser(p)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"status": 1,

			"message": err,
		})

	}
	res, err := database.DBSplash.Query("INSERT into splash(name, date) VALUES ($1, $2)", p.Name, p.Date)

	if err != nil {
		return err
	}
	return c.Status(200).JSON(&fiber.Map{"status": 0, "message": p, "resp": res})

}

func GetAllSplash(c *fiber.Ctx) error {

	rep, err := repository.GetSplash()

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 400, "error": err})
	}

	return c.Status(200).JSON(&fiber.Map{"statusCode": 200, "data": rep})
}
