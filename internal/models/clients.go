package models

import "time"

type Client struct {
	ID           uint
	UserID       uint
	BusinessName string
	Address      Address
	ABN          string
	CreatedDate  time.Time
	Deleted      bool
}
