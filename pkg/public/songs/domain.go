package songs

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Songs  Model struct Songs
type Songs struct {
	ID          string     `json:"id" db:"id" valid:"required,uuid"`
	Name        string     `json:"name" db:"name" valid:"required"`
	Artist      string     `json:"artist" db:"artist" valid:"required"`
	Album       string     `json:"album" db:"album" valid:"required"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleter *string    `json:"user_deleter" db:"user_deleter"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator string     `json:"user_creator" db:"user_creator"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

func NewSongs(id string, name string, artist string, album string) *Songs {
	return &Songs{
		ID:        id,
		Name:      name,
		Artist:    artist,
		Album:     album,
		IsDeleted: false,
	}
}

func (m *Songs) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
