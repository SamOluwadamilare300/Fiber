package main

import (
	"log"
	// "strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/user/:name/:id", func(c *fiber.Ctx) error {
        //Coventional way to get Params
		// name := c.Params("name")
		// id,  _ := strconv.Atoi(c.Params("id"))
		// id, _ := c.ParamsInt("id")

		//Another way to get Params
		type UserParams struct {
			Name string `params:"name"`
			ID   int64 `params:"id"`
		}

		userParams := new(UserParams) 

		if err := c.ParamsParser(userParams); err != nil{
			return err
		}
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"name": userParams.Name,
		"id": userParams.ID,
		})
	  })




	log.Fatal(app.Listen(":3000"))

}
