package repository

import (
	"github.com/Toolnado/SwalloW/model"
	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{
		db: db,
	}
}

func (r *UsersPostgres) GetUserByID(id int) (model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE id=$1"

	if err := r.db.Get(&user, query, id); err != nil {
		return user, err
	}

	return user, nil

}

func (r *UsersPostgres) GetAllUsers() ([]model.User, error) {
	var users []model.User

	query := "SELECT * FROM users"
	if err := r.db.Select(&users, query); err != nil {
		return users, err
	}

	return users, nil
}
