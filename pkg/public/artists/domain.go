package artists

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Artists  Model struct Artists
type Artists struct {
	ID          string     `json:"id" db:"id" valid:"required,uuid"`
	Name        string     `json:"name" db:"name" valid:"required"`
	Surname     string     `json:"surname" db:"surname" valid:"required"`
	Nationality string     `json:"nationality" db:"nationality" valid:"required"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleter *string    `json:"user_deleter" db:"user_deleter"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator string     `json:"user_creator" db:"user_creator"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

func NewArtists(id string, name string, surname string, nationality string) *Artists {
	return &Artists{
		ID:          id,
		Name:        name,
		Surname:     surname,
		Nationality: nationality,
		IsDeleted:   false,
	}
}

func (m *Artists) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
