package repository

import (
	"fmt"

	"github.com/Toolnado/SwalloW/model"
	"github.com/jmoiron/sqlx"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsPostgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{
		db: db,
	}
}

func (r *PostsPostgres) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post

	query := fmt.Sprintln("SELECT * FROM posts")
	if err := r.db.Select(&posts, query); err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *PostsPostgres) GetAllPostsThisUser(id int) ([]model.Post, error) {
	var posts []model.Post

	query := fmt.Sprintln("SELECT * FROM posts WHERE user_id=$1")
	if err := r.db.Select(&posts, query, id); err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *PostsPostgres) CreatePost(p *model.Post) (int, error) {
	var id int
	query := fmt.Sprintln("INSERT INTO posts (post_title, post_description, user_id) VALUES($1, $2, $3) RETURNING id")

	row := r.db.QueryRow(query, p.Title, p.Description, p.User)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *PostsPostgres) GetPost(id int) (model.Post, error) {
	var post model.Post
	query := fmt.Sprintln("SELECT * FROM posts WHERE id=$1")

	if err := r.db.Get(&post, query, id); err != nil {
		return post, err
	}

	return post, nil
}
