package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/login", func(c *fiber.Ctx) error {
		// Create cookie
		cookieAuth := new(fiber.Cookie)
		cookieAuth.Name = "username"
		cookieAuth.Value = "samdare"
		cookieAuth.Expires = time.Now().Add(24 * time.Hour)
		cookieAuth.MaxAge = 100

		cookieTheme := new(fiber.Cookie)
		cookieTheme.Name = "app_theme"
		cookieTheme.Value = "dark"



		// Set cookie
		
		c.Cookie(cookieAuth)
		c.Cookie(cookieTheme)

	     return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/checkout", func( c *fiber.Ctx) error{
        
		fmt.Println("username", c.Cookies("username", "No user found!"))
		fmt.Println("app_theme", c.Cookies("app_name", "Light found!"))
		
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/logout", func(c *fiber.Ctx) error {

		  // Clears all cookies:
		  c.ClearCookie()

	 return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/verify-cookie", func(c *fiber.Ctx) error{
		type AuthCookie struct {
			Username string `cookie:"username"`
			AppTheme string `cookie:"app_theme"`
		}

		auth := new(AuthCookie)

		if err := c.CookieParser(auth); err != nil {
			return err 
		}

		fmt.Println("Username", auth.Username)
		fmt.Println("AppThemE", auth.AppTheme)


		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(":3000"))
}