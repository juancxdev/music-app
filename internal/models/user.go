package models

import (
	"time"
)

// User  Model struct User
type Users struct {
	ID           string     `json:"id" db:"id" valid:"required,uuid"`
	Name         string     `json:"name" db:"name" valid:"required"`
	Email        string     `json:"email" db:"email" valid:"required"`
	CreationDate time.Time  `json:"creationDate" db:"creationDate" valid:"required"`
	IsDeleted    bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleter  *string    `json:"user_deleter" db:"user_deleter"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator  string     `json:"user_creator" db:"user_creator"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}
