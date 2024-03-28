package controllers

import (
	"context"
	"fmt"

	"github.com/HunterX18/book_my_ride/booking/db"
	"github.com/gofiber/fiber/v2"
)

type BookRideRequest struct {
	RideID int64 `json:"ride_id"`
	OwnerID int64 `json:"owner_id"`
}

func BookRide(c *fiber.Ctx) error {
		var bookRide BookRideRequest;
		if err := c.BodyParser(&bookRide); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		Db := db.Db;
		ctx := context.Background();
		txn, err := Db.BeginTx(ctx, nil);

		if err != nil {
			return c.Status(400).JSON(err.Error())
		}
		defer txn.Rollback();

		_, err = txn.Exec(db.UpdateRide, bookRide.OwnerID, "booked", bookRide.RideID);
		if err != nil {
			return c.Status(400).JSON(err.Error())
		}

		txn.Commit();

		return c.JSON(fmt.Sprintf("Ride %d booked successfully by %d", bookRide.RideID, bookRide.OwnerID));
}