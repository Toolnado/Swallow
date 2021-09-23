package service

import (
	"github.com/Toolnado/SwalloW/internal/repository"
	"github.com/Toolnado/SwalloW/model"
)

type Authentication interface {
	CreateUser(user *model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Users interface {
	GetUserByID(id int) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetMyAccount(id int) (model.User, error)
}

type Posts interface {
	GetAllPosts() ([]model.Post, error)
	GetAllPostsThisUser(id int) ([]model.Post, error)
	CreatePost(p *model.Post) (int, error)
	GetPost(id int) (model.Post, error)
}

type Service struct {
	Authentication Authentication
	Users          Users
	Posts          Posts
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repo.Authentication),
		Users:          NewUsersService(repo.Users),
		Posts:          NewPostsService(repo.Posts),
	}
}
