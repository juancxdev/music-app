package albums

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ResponseAlbums struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type RequestAlbums struct {
	ID          string    `json:"id" valid:"required,uuid"`
	Name        string    `json:"name"  valid:"required"`
	Artist      string    `json:"artist"  valid:"required"`
	ReleaseDate time.Time `json:"releaseDate"  valid:"required"`
}

func (m *RequestAlbums) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
