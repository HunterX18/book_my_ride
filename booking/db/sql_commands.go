package db

var CreateRide string = `INSERT INTO rides (Owner_ID, Status) VALUES (1, 'free')`

var SelectRide string = `SELECT ID FROM rides WHERE Status = 'free' LIMIT 1 FOR UPDATE`

var SelectRideX string = `SELECT ID FROM rides WHERE ID = $1 FOR UPDATE`

var UpdateRide string = `UPDATE rides SET Owner_ID = $1, Status = $2 WHERE ID = $3`