package services

import "github.com/sawmeraw/ims/internal/database"

type Services struct {
	// User UserService
}

func NewServices(db *database.Database) *Services {
	return &Services{
		// User: NewUserService(db),
	}
}
