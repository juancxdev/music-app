package albums

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Albums  Model struct Albums
type Albums struct {
	ID          string     `json:"id" db:"id" valid:"required,uuid"`
	Name        string     `json:"name" db:"name" valid:"required"`
	Artist      string     `json:"artist" db:"artist" valid:"required"`
	ReleaseDate time.Time  `json:"releaseDate" db:"releasedate" valid:"required"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleter *string    `json:"user_deleter" db:"user_deleter"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator string     `json:"user_creator" db:"user_creator"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

func NewAlbums(id string, name string, artist string, releaseDate time.Time) *Albums {
	return &Albums{
		ID:          id,
		Name:        name,
		Artist:      artist,
		ReleaseDate: releaseDate,
		IsDeleted:   false,
	}
}

func (m *Albums) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
