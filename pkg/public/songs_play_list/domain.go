package songs_play_list

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// SongsOfPlaylist  Model struct SongsOfPlaylist
type SongsOfPlaylist struct {
	ID          string     `json:"id" db:"id" valid:"required,uuid"`
	Playlist    int        `json:"playlist" db:"playlist" valid:"required"`
	Song        int        `json:"song" db:"song" valid:"required"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleter *string    `json:"user_deleter" db:"user_deleter"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator string     `json:"user_creator" db:"user_creator"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

func NewSongsOfPlaylist(id string, playlist int, song int) *SongsOfPlaylist {
	return &SongsOfPlaylist{
		ID:        id,
		Playlist:  playlist,
		Song:      song,
		IsDeleted: false,
	}
}

func (m *SongsOfPlaylist) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
