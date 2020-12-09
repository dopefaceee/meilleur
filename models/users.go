package models

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// ErrNotFound string
	ErrNotFound = errors.New("models: resource not found")

	// ErrInvalidID for > 0
	ErrInvalidID = errors.New("models: ID must be greater than 0")
)

// UserService struct
type UserService struct {
	db *gorm.DB
}

// DestructiveReset wow
func (us *UserService) DestructiveReset() {
	us.db.Exec("DROP TABLE users")
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

// Create will create user and backfill data
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Update user in db
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete user in db
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}

	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// ByID , If the user is found, return nil error
// if the user is not found, return ErrNotFound
// If there is another error , returh error with information
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := db.First(&user).Error
	return &user, err
}

// ByEmail for query with email
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// first will query using provided db
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}

	return err

}

// AutoMigrate helper
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}); err != nil {
		return err
	}

	return nil
}

// User model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
