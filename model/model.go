package model

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username" validate:"required, min=2, max=20"`
	Password  string    `json:"password" db:"password" validate:"required, min=6, max=100"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	CreateAt  time.Time `json:"create_at" db:"create_at"`
}

type Post struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"post_title" db:"post_title"`
	Description string `json:"post_description" db:"post_description"`
	User        int    `josn:"user_id" db:"user_id"`
}
