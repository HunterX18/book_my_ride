package main

import (
	"github.com/HunterX18/go_txn/controllers"
	"github.com/HunterX18/go_txn/db/sqlx"
	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()

	sqlx.InitDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello World")
	});

	app.Post("/accounts", controllers.CreateAccount)
	app.Post("/transfer", controllers.TransferMoney)
	app.Post("/get-balance", controllers.GetBalance)


	app.Listen(":5000")
}