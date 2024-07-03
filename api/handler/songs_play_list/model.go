package songs_play_list

import (
	"github.com/asaskevich/govalidator"
)

type ResponseSongsOfPlaylist struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type RequestSongsOfPlaylist struct {
	ID       string `json:"id" valid:"required,uuid"`
	Playlist int    `json:"playlist"  valid:"required"`
	Song     int    `json:"song"  valid:"required"`
}

func (m *RequestSongsOfPlaylist) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
