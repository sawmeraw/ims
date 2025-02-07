package models

import "time"

type User struct {
	ID            int
	FirstName     string
	LastName      string
	Address       Address
	ContactEmail  string
	ContactNumber string
	ABN           string
	BankAccount   BankAccount
	Clients       []Client
	Login         Login
	CreatedDate   time.Time
	Deleted       bool
}

type Login struct {
	email    string
	password string
}

type Address struct {
	Street   string
	City     string
	PostCode string
}

type BankAccount struct {
	FirstName     string
	LastName      string
	BSB           string
	AccountNumber string
}
