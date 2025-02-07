package models

import "time"

type Client struct {
	ID           int
	UserID       int
	BusinessName string
	Address      Address
	ABN          string
	CreatedDate  time.Time
	Deleted      bool
}
