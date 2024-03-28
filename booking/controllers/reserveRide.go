package controllers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/HunterX18/book_my_ride/booking/db"
	"github.com/gofiber/fiber/v2"
)

type ReserveRideRequest struct {
	Owner_ID int64 `json:"owner_id"`
}

func ReserveRide(c *fiber.Ctx) error {
	var reserveRideRequest ReserveRideRequest;
	if err := c.BodyParser(&reserveRideRequest); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	Db := db.Db;
	txn, err := Db.BeginTx(context.Background(), nil);
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	defer txn.Rollback();

	var reservedRide int64;
	res := txn.QueryRow(db.SelectRide);
	if res.Err() != nil {
		return c.Status(400).JSON(res.Err().Error())
	}
	err = res.Scan(&reservedRide);
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(400).JSON("No rides are available for booking right now")
		}
		return c.Status(400).JSON(err.Error());
	}
	
	_, err = txn.Exec(db.UpdateRide, reserveRideRequest.Owner_ID, "reserved", reservedRide);
	if err != nil {
		fmt.Println("error while reserving")
		return c.Status(400).JSON(err.Error());
	}

	txn.Commit();
	return c.Status(200).JSON(reservedRide)
}