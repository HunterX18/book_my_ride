package main

import (
	"github.com/HunterX18/book-my-ride/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Order server is running")
	})

	app.Post("/place-order", controllers.PlaceOrder);

	app.Listen(":5002")

}