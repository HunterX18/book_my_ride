package controllers

import (
	"context"

	"github.com/HunterX18/go_txn/db/sqlx"
	"github.com/gofiber/fiber/v2"
)

type GetBalanceRequest struct {
	OwnerID int64 `json:"owner_id"`
}


func GetBalance(c *fiber.Ctx) error {
		var getBalanceRequest GetBalanceRequest;
		if err := c.BodyParser(&getBalanceRequest); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		db := sqlx.Db
		ctx := context.Background();
		txn, err := db.BeginTx(ctx, nil);

		if err != nil {
			return c.Status(400).JSON(err.Error())
		}

		defer txn.Rollback();

		getAccountBalance := `SELECT Balance FROM accounts WHERE ID = $1;`

		res := txn.QueryRow( getAccountBalance, getBalanceRequest.OwnerID)

		if res.Err() != nil {
			return c.Status(400).JSON(res.Err().Error())
		}
		var accountBalance int64;
		err = res.Scan(&accountBalance)
		if err != nil {
			c.Status(400).JSON(err.Error())
		}
		txn.Commit();
		return c.JSON(accountBalance);
}