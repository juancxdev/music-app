package play_lists

import (
	"github.com/asaskevich/govalidator"
)

type ResponsePlaylists struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type RequestPlaylists struct {
	ID   string `json:"id" valid:"required,uuid"`
	Name string `json:"name"  valid:"required"`
	User int    `json:"user"  valid:"required"`
}

func (m *RequestPlaylists) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
