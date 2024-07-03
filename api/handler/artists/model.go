package artists

import (
	"github.com/asaskevich/govalidator"
)

type ResponseArtists struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type RequestArtists struct {
	ID          string `json:"id" valid:"required,uuid"`
	Name        string `json:"name"  valid:"required"`
	Surname     string `json:"surname"  valid:"required"`
	Nationality string `json:"nationality"  valid:"required"`
}

func (m *RequestArtists) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
