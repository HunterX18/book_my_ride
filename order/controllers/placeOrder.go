package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type PlaceOrderRequest struct {
	Owner_ID int64 `json:"owner_id"`
}

type GetBalanceRequest struct {
	Owner_ID int64 `json:"owner_id"`
}


type ReserveRideRequest struct {
	Owner_ID int64 `json:"owner_id"`
}

type GetBalanceResponse struct {
	Balance int64 `json:"balance"`
}

type ConfirmRideRequest struct {
	OwnerId int64 `json:"owner_id"`
	RideId int64 `json:"ride_id"`
}

type ConfirmPaymentRequest struct {
	ToAccountID int64 `json:"to_account_id"`
	FromAccountID int64 `json:"from_account_id"`
	Amount int64 `json:"amount"`
}

func PlaceOrder(c *fiber.Ctx) error {
	var placeOrderRequest PlaceOrderRequest;

	if err := c.BodyParser(&placeOrderRequest); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var getBalanceRequest GetBalanceRequest;
	getBalanceRequest.Owner_ID = placeOrderRequest.Owner_ID 

	agent := fiber.Post("http://localhost:5000/get-balance/");
	agent.JSON(getBalanceRequest)

    statusCode, body, err := agent.Bytes();

	if err != nil || statusCode != 200 {
		return c.Status(400).JSON("Couldn't get balance to continue transaction")
	}

	// check if balance is sufficient
	// var getBalanceResponse GetBalanceResponse;
	var balance int64
	if err := json.Unmarshal(body, &balance); err != nil {
		return c.Status(400).JSON(err.Error());
	}

	if balance < 100 {
		return c.Status(200).JSON("Insufficient Balance")
	}

	var reserveRideResponse int64;
	var reserveRideRequest ReserveRideRequest;
	reserveRideRequest.Owner_ID = placeOrderRequest.Owner_ID;

	// check if a ride is available
	agent = fiber.Post("http://localhost:5001/reserve-ride");
	agent.JSON(reserveRideRequest);
	statusCode, body, err = agent.Bytes();

	if err != nil || statusCode != 200 {
		return c.Status(400).JSON("Unable to reserve ride")
	}

	if err := json.Unmarshal(body, &reserveRideResponse); err != nil {
		return c.Status(400).JSON("Unable to parse reserveRideResponse")
	}


	// confirm ride
	var confirmRide ConfirmRideRequest;
	confirmRide.OwnerId = placeOrderRequest.Owner_ID;
	confirmRide.RideId = reserveRideResponse;
	agent = fiber.Post("http://localhost:5001/book-ride");
	agent.JSON(confirmRide);

	statusCode, _, err = agent.Bytes();

	if err != nil || statusCode != 200 {
		return c.Status(400).JSON("Unable to confirm booking")
	}


	// confirm payment
	var confirmPayment ConfirmPaymentRequest;
	confirmPayment.FromAccountID = placeOrderRequest.Owner_ID;
	confirmPayment.ToAccountID = 3;
	confirmPayment.Amount = 100

	agent = fiber.Post("http://localhost:5000/transfer");
	agent.JSON(confirmPayment)

	statusCode, _, err = agent.Bytes();

	if err != nil || statusCode != 200 {
		// TODO: Free the booked ride
		return c.Status(400).JSON("Unable to transfer money")
	}

	return c.Status(200).JSON(fmt.Sprintf("Ride %d booked successfully by %d", reserveRideResponse, placeOrderRequest.Owner_ID));
}