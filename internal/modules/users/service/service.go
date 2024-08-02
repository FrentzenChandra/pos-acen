package service

import (
	"pos-acen/internal/modules/users/ports"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

// func (s *userService) RegisterUser(bReq entity.User) (*uuid.UUID, error) {

// 	// id, err := s.repo.RegisterUser(bReq)
// 	// if err != nil {
// 	// 	log.Println("Error registering user ")
// 	// 	return nil, err
// 	// }
// 	// return id , err
// }
