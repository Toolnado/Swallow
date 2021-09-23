package service

import (
	"log"

	"github.com/Toolnado/SwalloW/internal/repository"
	"github.com/Toolnado/SwalloW/model"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(r repository.Users) *UsersService {
	return &UsersService{
		repo: r,
	}
}

func (s *UsersService) GetUserByID(id int) (model.User, error) {
	user, err := s.repo.GetUserByID(id)

	if err != nil {
		log.Printf("[user not found: %s]\n", err)
		return user, err
	}

	user.Password = "HIDDEN"

	return user, nil
}

func (s *UsersService) GetAllUsers() ([]model.User, error) {

	users, err := s.repo.GetAllUsers()

	if err != nil {
		log.Printf("[users not found: %s]\n", err)
		return users, err
	}

	for i := 0; i < len(users); i++ {
		users[i].Password = "HIDDEN"
	}

	return users, nil
}

func (s UsersService) GetMyAccount(id int) (model.User, error) {
	user, err := s.repo.GetUserByID(id)

	if err != nil {
		log.Printf("[users not found: %s]\n", err)
		return user, err
	}

	return user, err
}
