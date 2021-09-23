package repository

import (
	"github.com/Toolnado/SwalloW/model"
	"github.com/jmoiron/sqlx"
)

type Authentication interface {
	CreateUser(u *model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Users interface {
	GetUserByID(id int) (model.User, error)
	GetAllUsers() ([]model.User, error)
}

type Posts interface {
	GetAllPosts() ([]model.Post, error)
	GetAllPostsThisUser(id int) ([]model.Post, error)
	CreatePost(p *model.Post) (int, error)
	GetPost(id int) (model.Post, error)
}

type Repository struct {
	Authentication Authentication
	Users          Users
	Posts          Posts
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authentication: NewAuthPostgres(db),
		Users:          NewUsersPostgres(db),
		Posts:          NewPostsPostgres(db),
	}
}
