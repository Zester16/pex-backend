package controller

//"strconv"
import (
	"fmt"
	"time"

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
			"status":  1,
			"message": err,
		})

	}
	if p.Name == "" {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 1, "statusMessage": "Kindly add name of newspaper"})
	}
	if p.Image_Url == "" {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 1, "statusMessage": "Kindly add image url of newspaper"})
	}
	if p.Epaper_Url == "" {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 1, "statusMessage": "epaper url should not be empty"})
	}

	p.Id = uuid.New().String()
	currentTime := time.Now().UTC().Unix()
	p.Created_At = &currentTime
	err = repository.AddNewspaper(*p)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 400, "statusMessage": err.Error()})
	}
	return c.JSON(&fiber.Map{"statusCode": 0, "statusMessage": "success"})
}

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
	newsreadList, err := repository.GetNewsreadwithDateAndName(p.Newspaper_Id, p.Read_At)

	if len(newsreadList) > 0 || err != nil {
		//fmt.Println(newsreadList, err.Error())
		return c.Status(400).JSON(&fiber.Map{"statusCode": 1, "statusMessage": "Duplicate Data Exists"})
	}
	err = repository.AddNewsRead(*p)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 400, "statusMessage": err.Error()})
	}
	return c.JSON(&fiber.Map{"statusCode": 0, "statusMessage": "success"})
}

// ****READ NEWS
// get all news letters in order as most read
func GetNewspapersPaginated(c *fiber.Ctx) error {

	id := c.Query("id")
	fmt.Println("id", id)
	resp, err := repository.GetNewspaperPaginated(id)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": "1", "statusMessage": err.Error()})
	}
	total, err := repository.GetNewspaperTotalCount()

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 1, "statusMessage": err})
	}
	return c.JSON(&fiber.Map{"statusCode": 0, "data": resp, "total": total})
}

func GetAllNewspapers(c *fiber.Ctx) error {
	allNewspapers, err := repository.GetNewspapersAll()

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": "1", "statusMessage": err.Error()})
	}
	return c.JSON(&fiber.Map{"statusCode": 0, "data": allNewspapers})
}

func GetNewsReadAll(c *fiber.Ctx) error {
	newsreadAll, err := repository.GetNewsReadAll()

	if err != nil {

		return c.Status(400).JSON(&fiber.Map{"statusCode": "1", "statusMessage": err.Error()})
	}

	return c.JSON(&fiber.Map{"statusCode": 0, "data": newsreadAll})
}

// ********HELPER FUNCTION FOR checking new related errorß
func checkNewspaperInputFromRequestBody() bool {

	return true
}
