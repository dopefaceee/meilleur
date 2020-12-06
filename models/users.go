package models

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// ErrNotFound string
	ErrNotFound = errors.New("models: resource not found")
)

// UserService struct
type UserService struct {
	db *gorm.DB
}

// DestructiveReset wow
func (us *UserService) DestructiveReset() {
	//us.db.DropTableifExists(&User{})
	us.db.AutoMigrate(&User{})
}

// NewUserService for open conn
func NewUserService(connectionInfo string) (*UserService, error) {

	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return &UserService{
		db: db,
	}, nil
}

// ByID , If the user is found, return nil error
// if the user is not found, return ErrNotFound
// If there is another error , returh error with information
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// User model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
