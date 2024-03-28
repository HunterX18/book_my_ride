package db

import "time"

type Ride struct {
	ID int64 `json:"id"`
	OwnerID string `json:"owner"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
