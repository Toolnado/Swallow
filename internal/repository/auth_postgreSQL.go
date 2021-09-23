package repository

import (
	"github.com/Toolnado/SwalloW/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(u *model.User) (int, error) {
	var id int
	query := "INSERT INTO users (username, password, first_name, last_name, create_at) VALUES($1, $2, $3, $4, $5) RETURNING id"
	row := r.db.QueryRow(query, u.Username, u.Password, u.FirstName, u.LastName, u.CreateAt)

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := "SELECT id FROM users WHERE username=$1 AND password=$2"

	if err := r.db.Get(&user, query, username, password); err != nil {
		return user, err
	}

	return user, nil
}
