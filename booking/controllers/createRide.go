package controllers

import (
	"context"

	"github.com/HunterX18/book_my_ride/booking/db"
	"github.com/gofiber/fiber/v2"
)

func CreateRide(c *fiber.Ctx) error {
	Db := db.Db;
	txn, err := Db.BeginTx(context.Background(), nil);
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	defer txn.Rollback();
	_, err = txn.Exec(db.CreateRide);
	if err != nil {
		c.Status(400).JSON(err.Error())
	}


	txn.Commit();
	return c.JSON("Ride created successfully");
}