package services

// import (
// 	"github.com/sawmeraw/ims/internal/database"
// 	"github.com/sawmeraw/ims/internal/models"
// )

// type UserService interface {
// 	Create(user models.User) (uint, error)
// 	GetById(id uint) (*models.User, error)
// }

// type userService struct {
// 	db *database.Database

// func NewUserService(db *database.Database) UserService {
// 	return &userService{db: db}
// }

// func (s *userService) Create(user models.User) (uint, error) {
// 	return user.ID, nil
// }

// // func (s *userService) GetById(id uint) (*models.User, error) {
// // 	var user models.User
// // 	return s.db.First(&user, id).Error
// // }
