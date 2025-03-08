package main

import (

	"log"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	// app.Get("/store/:name/:books", func(c *fiber.Ctx) error{
	// 	params := c.AllParams()
    //     // fmt.Println(c.AllParams(), "Book one is avaliable")
		
	// 	fmt.Println(params["name"])
	// 	fmt.Println(params["books"])


	// 	return c.SendStatus(fiber.StatusOK)
	// })


	//Download FILE Function

	// app.Get("/", func(c *fiber.Ctx) error{
	// 	return c.Download("./files/Sam O. Afolabi CV.pdf")
	// })



	// Field names should start with an uppercase letter
type Person struct {
    Name string `json:"name" `
    Pass string `json:"pass" `
}

app.Post("/", func(c *fiber.Ctx) error {
        p := new(Person)

		NewId := 123454

        if err := c.BodyParser(p); err != nil {
            return err
        }
		

        log.Println(p.Name) // john
        log.Println(p.Pass) // doe

       return c.Status(fiber.StatusOK).JSON(fiber.Map{
       "name": p.Name,
	   "pass": p.Pass,
	   "NewId": NewId,
	   })
})

// Run tests with the following curl commands

// curl -X POST -H "Content-Type: application/json" --data "{\"name\":\"john\",\"pass\":\"doe\"}" localhost:3000

// curl -X POST -H "Content-Type: application/xml" --data "<login><name>john</name><pass>doe</pass></login>" localhost:3000

// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data "name=john&pass=doe" localhost:3000

// curl -X POST -F name=john -F pass=doe http://localhost:3000

// curl -X POST "http://localhost:3000/?name=john&pass=doe"



	log.Fatal(app.Listen(":3000"))
}