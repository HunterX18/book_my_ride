package controllers

import (
	"context"

	"github.com/HunterX18/go_txn/db"
	"github.com/HunterX18/go_txn/db/sqlx"
	"github.com/gofiber/fiber/v2"
)

type TransferRequest struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

type Account struct {
	ID int64
	Owner string
	Balance int64
}

func TransferMoney(c *fiber.Ctx) error {

	var transferRequest TransferRequest
	if reqErr := c.BodyParser(&transferRequest); reqErr != nil {
		return c.Status(400).JSON(reqErr.Error())
	}

	// ensures queries are executed in the same order to avoid deadlocks
	var fromAccountID, toAccountID, amount int64;
	if(transferRequest.FromAccountID < transferRequest.ToAccountID) {
		fromAccountID = transferRequest.FromAccountID;
		toAccountID = transferRequest.ToAccountID;
		amount = transferRequest.Amount 
	} else {
		fromAccountID = transferRequest.ToAccountID;
		toAccountID = transferRequest.FromAccountID
		amount = -transferRequest.Amount
	}

	Db := sqlx.Db;

	// beginning transaction
	txn, txnErr := Db.BeginTx(context.Background(), nil);
	if txnErr != nil {
		return c.Status(400).JSON(txnErr.Error())
	}
	defer txn.Rollback()

	// creating a transfer record
	_, trfErr := txn.Exec(db.CreateTransfer, fromAccountID, toAccountID, amount)
	if trfErr != nil {
		return c.Status(400).JSON(trfErr.Error())
	}

	// acquiring row level lock to update balance of fromAccount
	_, err := txn.Exec(db.GetAccountQuery, fromAccountID)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	// updating balance of fromAccount
	_, err = txn.Exec(db.UpdateAccountQuery, -amount, fromAccountID)
	if err != nil {
		c.JSON(err.Error())
	}

	// acquiring row level lock to update balance of toAccount
	_, err = txn.Exec(db.GetAccountQuery, toAccountID);
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// updating balance of toAccount
	_, err = txn.Exec(db.UpdateAccountQuery, amount, toAccountID)
	if err != nil {
		c.Status(400).JSON(err.Error())
	}

	// commiting transaction
	txn.Commit();
	return c.JSON("Transfer successful")
}