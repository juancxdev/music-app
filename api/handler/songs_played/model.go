package songs_played

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ResponseSongsPlayed struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type RequestSongsPlayed struct {
	ID   string    `json:"id" valid:"required,uuid"`
	User int       `json:"user"  valid:"required"`
	Song int       `json:"song"  valid:"required"`
	Date time.Time `json:"date"  valid:"required"`
}

func (m *RequestSongsPlayed) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
