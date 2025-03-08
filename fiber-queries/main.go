package main

import (
	"log"
	// "strconv"

	"github.com/gofiber/fiber/v2"
)

func main(){

	app := fiber.New()

	// GET  http://127.0.0.1:3000/?sort=asc&date=today&new=true&maxprice=100.66

app.Get("/", func(c *fiber.Ctx) error {
	
	// queries := c.Queries()

// 	sort := c.Query("sort")
// 	date := c.QueryInt("date")
// 	new  := c.QueryBool("new")
// 	maxprice := c.QueryFloat("maxprice")

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"sort": sort,
// 		"date": date,
// 		"new": new,
// 		"maxprice" : maxprice,
// 	})
// })


//Queries Parser

type QueryProduct struct {
	Sort string `query:"sort"`
	Date int `query:"date"`
	New bool `query:"new"`
	MaxPrice float64 `query:"maxprice"`
}

    queryProduct := new(QueryProduct)

if err := c.QueryParser(queryProduct); err != nil{
	return err 
}

return c.Status(fiber.StatusOK).JSON(fiber.Map{
	"sort": queryProduct.Sort,
	"date": queryProduct.Date,
	"new": queryProduct.New,
	"maxprice" : queryProduct.MaxPrice,
})
})




log.Fatal(app.Listen(":3000"))
}