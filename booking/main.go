package main

import (
	"github.com/HunterX18/book_my_ride/booking/controllers"
	"github.com/HunterX18/book_my_ride/booking/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();

	db.InitDB();
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Booking server is running")
	})

	app.Get("/create-ride", controllers.CreateRide);
	app.Post("/book-ride", controllers.BookRide);
	app.Post("/reserve-ride", controllers.ReserveRide)

	app.Listen(":5001")

}