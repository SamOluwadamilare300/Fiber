package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)
  

func main(){
	app :=  InitApp()
	log.Fatal(app.Listen(":3000"))
}


func InitApp()  *fiber.App {
	app := fiber.New()

	// appV1 := app.Group("/v1")
	// appV1.Get("/", func(c *fiber.Ctx) error{
	// 	return c.SendString("Version one")
	// })
	// appV2 := app.Group("/v2")
	// appV2.Get("/", func(c *fiber.Ctx) error{
	// 	return c.SendString("Version two")
	// })

	// app.Route("/user", func(api fiber.Router) {
	// 	api.Get("/", func(c *fiber.Ctx) error{return c.SendString("User Info")})
	// 	api.Post("/", func(c *fiber.Ctx) error{return c.SendString("User Creat Account")})
	// 	api.Get("/:id", func(c *fiber.Ctx) error{return c.SendString("User fetch info by Id")})
	// 	api.Patch("/:id", func(c *fiber.Ctx) error{return c.SendString("User modify info")})
	// 	api.Delete("/:id",func(c *fiber.Ctx) error{return c.SendString("User delete Info")})

	// })

	/// Sever Shutdown

	// app.Get("/api", func(c *fiber.Ctx) error{
	// 	if c.IP() != "127.0.0.1"{
	// 		return app.Shutdown()
	// 	}

	// 	//    Use this as wrong IP to shutdown server manually: 227.3.1.2

	// 	return c.SendString("Correct IP Address connected!")
	// })

	app.Get("/check-token", func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// Here you would typically validate the token
		// For simplicity, we'll just check if it's "Bearer mock-token"
		if token != "Bearer mock-token" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusOK)
	})

	return app
}

