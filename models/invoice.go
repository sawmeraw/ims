package models

import "time"

type Invoice struct {
	ID          int
	UserID      int
	ClientID    int
	Client      Client
	User        User
	Items       []InvoiceItem
	CreatedDate time.Time
	SentDate    time.Time
	TotalAmount float64
	Status      string
	Recurring   bool
}

type InvoiceItem struct {
	ID          int
	InvoiceID   int
	Description string
	Quantity    int
	UnitPrice   float64
	TotalPrice  float64
}

type Status int

const (
	Pending Status = iota
	Paid
	Overdue
	Cancelled
)

func (s Status) String() string {
	return [...]string{"Pending", "Paid", "Overdue", "Cancelled"}[s]
}
