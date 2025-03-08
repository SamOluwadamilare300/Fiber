
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	//SEND

    //  app.Get("/:name", func(c *fiber.Ctx) error{
	// 	name := []byte(c.Params("name"))

	// 	return c.Send(name)
	//  })

	//SEND-FILE


	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendFile("./files/dummy.txt");
	  
		// Disable compression
		// return c.SendFile("./static/index.html", false);
	  //})


       //STATUS
	  app.Get("/fiber", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		return nil
	  })
	  
	  app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	  })
	  
	  app.Get("/world", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./public/index.html")
	  })


	log.Fatal(app.Listen(":3000"))
}