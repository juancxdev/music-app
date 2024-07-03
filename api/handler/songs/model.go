package songs

import (
	"github.com/asaskevich/govalidator"
)

type ResponseSongs struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type RequestSongs struct {
	ID     string `json:"id" valid:"required,uuid"`
	Name   string `json:"name"  valid:"required"`
	Artist string `json:"artist"  valid:"required"`
	Album  string `json:"album"  valid:"required"`
}

func (m *RequestSongs) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
