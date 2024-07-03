package songs_played

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// SongsPlayed  Model struct SongsPlayed
type SongsPlayed struct {
	ID          string     `json:"id" db:"id" valid:"required,uuid"`
	User        int        `json:"user" db:"user" valid:"required"`
	Song        int        `json:"song" db:"song" valid:"required"`
	Date        time.Time  `json:"date" db:"date" valid:"required"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleter *string    `json:"user_deleter" db:"user_deleter"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator string     `json:"user_creator" db:"user_creator"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

func NewSongsPlayed(id string, user int, song int, date time.Time) *SongsPlayed {
	return &SongsPlayed{
		ID:        id,
		User:      user,
		Song:      song,
		Date:      date,
		IsDeleted: false,
	}
}

func (m *SongsPlayed) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
