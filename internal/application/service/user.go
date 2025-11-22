package service

import (
	"THackaton/domain/repository"
	"THackaton/domain/service"
	"errors"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) service.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(req service.RegisterRequest) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("username and password are required")
	}
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	exists, err := s.userRepo.ExistsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already taken")
	}

	return s.userRepo.Create(req.Username, req.Email, req.Password)
}
