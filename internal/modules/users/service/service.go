package service

import (
	"log"
	"pos-acen/internal/modules/users/entity"
	"pos-acen/internal/modules/users/ports"

	"github.com/google/uuid"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(bReq entity.User) (*uuid.UUID, error) {
	user, err := s.repo.GetUserDetails(bReq)

	if err != nil {
		log.Println("Error registering user GETUserDetails Err : " + err.Error())
		return nil, err
	}

	if user != nil {
		log.Println("User email is already registered")
		return nil, err
	}

	id, err := s.repo.RegisterUser(bReq)

	if err != nil {
		log.Println("Error registering user RegisterUser Err : " + err.Error())
		return nil, err
	}

	return id, nil
}
