package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	type keyType struct{}
	var userKey keyType
	
	app.Use(func(c *fiber.Ctx) error {
	  c.Locals(userKey, "admin") // Stores the string "admin" under a non-exported type key
	  return c.Next()
	})
	
	app.Get("/admin", func(c *fiber.Ctx) error {
	  user, ok := c.Locals(userKey).(string) // Retrieves the data stored under the key and performs a type assertion
	  if ok && user == "admin" {
		return c.Status(fiber.StatusOK).SendString("Welcome, admin!")
	  }
	  return c.SendStatus(fiber.StatusForbidden)
	})

	log.Fatal(app.Listen(":3000"))
}