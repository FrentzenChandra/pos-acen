package entity

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt string    `json:"deleted_at"`
}
