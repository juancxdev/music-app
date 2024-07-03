package play_lists

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Playlists  Model struct Playlists
type Playlists struct {
	ID          string     `json:"id" db:"id" valid:"required,uuid"`
	Name        string     `json:"name" db:"name" valid:"required"`
	User        int        `json:"user" db:"user" valid:"required"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleter *string    `json:"user_deleter" db:"user_deleter"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator string     `json:"user_creator" db:"user_creator"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

func NewPlaylists(id string, name string, user int) *Playlists {
	return &Playlists{
		ID:        id,
		Name:      name,
		User:      user,
		IsDeleted: false,
	}
}

func (m *Playlists) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
