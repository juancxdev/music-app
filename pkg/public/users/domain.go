package users

import (
	"music-app/internal/models"
	"time"

	"github.com/asaskevich/govalidator"
)

// Users  Model struct Users
type Users models.Users

func NewUsers(id string, name string, email string, creationDate time.Time) *Users {
	return &Users{
		ID:           id,
		Name:         name,
		Email:        email,
		CreationDate: creationDate,
		IsDeleted:    false,
	}
}

func (m *Users) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
