package main

import (
	"fmt"
	"log"
	"time"

	// "time"

	"github.com/gofiber/fiber/v2"
)


func main(){
      app := fiber.New()

	// app.Post("/login", func(c *fiber.Ctx) error {
	// 	// Get first value from form field "name":
	// 	username := c.FormValue("username")
	// 	password := c.FormValue("password")


	// 	if username != "samdare" || password != "samdare123"{
	// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 			"message": "You're not authorised",
	// 		})
	// 	}
		
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	 "message": "Welcome",

	// 	})
	//   })


	//METHOD TO UPLAOD ONE DOCUMENT

	// app.Post("/upload-document", func(c *fiber.Ctx) error {
	// 	// Get first file from form field "document":
	// 	file, _ := c.FormFile("document")
	  
	// 	// Save file to root directory:
	// 	return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
	//   })


	  //METHOD TO UPLAOD MULTIPLE DOCUMENT

	  app.Post("/upload-document", func(c *fiber.Ctx) error{
		if form, err := c.MultipartForm(); err == nil{
			files :=  form.File["Documents"]
			for _, file := range files{
			  if err := c.SaveFile(file, fmt.Sprintf("./%s", time.Now().String())); err != nil {
				return err
			  }
			}
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	  })

log.Fatal(app.Listen(":3000"))
}