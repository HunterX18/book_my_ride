package controllers

import (
	"context"

	"github.com/HunterX18/go_txn/db/sqlx"
	"github.com/gofiber/fiber/v2"
)


type AccountRequest struct {
	Owner string `json:"owner"`
	Balance int64 `json:"balance"`
}

type ErrorMessage struct {
	Status int `json:"status"`
	Message string `json:"error"`
}

func CreateAccount(c *fiber.Ctx) error {
		var account AccountRequest;
		if err := c.BodyParser(&account); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		db := sqlx.Db
		ctx := context.Background();
		txn, err := db.BeginTx(ctx, nil);

		if err != nil {
			return c.Status(400).JSON(err.Error())
		}

		defer txn.Rollback();

		createAccount := `INSERT INTO accounts (Owner, Balance) VALUES ($1, $2) returning *;`

		_, err = txn.Exec(createAccount, account.Owner, account.Balance)

		if err != nil {
			return c.Status(400).JSON(err)
		}
		txn.Commit();
		return c.JSON("Account created successfully");
}