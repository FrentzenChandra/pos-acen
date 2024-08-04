package ports

import (
	"pos-acen/internal/modules/users/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	RegisterUser(bReq entity.User) (*uuid.UUID, error)
	GetUserDetails(bReq entity.User) (*entity.User, error)
}

type UserService interface {
	RegisterUser(bReq entity.User) (*uuid.UUID, error)
}
