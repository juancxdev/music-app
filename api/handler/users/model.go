package users

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ResponseUsers struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type RequestUsers struct {
	ID           string    `json:"id" valid:"required,uuid"`
	Name         string    `json:"name"  valid:"required"`
	Email        string    `json:"email"  valid:"required"`
	CreationDate time.Time `json:"creationDate"  valid:"required"`
}

func (m *RequestUsers) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
