package db

var CreateTransfer string = `INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES ($1, $2, $3)`

var CreateEntry string = `INSERT INTO entries (account_id, amount) VALUES ($1, $2)`

// lock the row while updating
var GetAccountQuery string = `SELECT id, balance FROM accounts WHERE id = $1 FOR NO KEY UPDATE`

var UpdateAccountQuery string = `UPDATE accounts SET balance = balance + $1 WHERE id = $2`