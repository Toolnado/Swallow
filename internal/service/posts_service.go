package service

import (
	"fmt"
	"log"

	"github.com/Toolnado/SwalloW/internal/repository"
	"github.com/Toolnado/SwalloW/model"
)

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(r repository.Posts) *PostsService {
	return &PostsService{
		repo: r,
	}
}

func (s *PostsService) GetAllPosts() ([]model.Post, error) {
	posts, err := s.repo.GetAllPosts()

	if err != nil {
		log.Printf("[posts not found: %s]\n", err)
		return posts, err
	}

	return posts, nil
}
func (s *PostsService) GetAllPostsThisUser(id int) ([]model.Post, error) {
	posts, err := s.repo.GetAllPostsThisUser(id)

	if err != nil {
		log.Printf("[posts this user not found: %s]\n", err)
		return posts, err
	}

	return posts, nil
}

func (s *PostsService) CreatePost(p *model.Post) (int, error) {
	id, err := s.repo.CreatePost(p)

	if err != nil {
		fmt.Printf("[not found id, error create user: %s]\n", err)
		return id, err
	}

	return id, err
}

func (s *PostsService) GetPost(id int) (model.Post, error) {
	post, err := s.repo.GetPost(id)

	if err != nil {
		fmt.Printf("[post not found: %s]\n", err)
		return post, err
	}

	return post, nil
}
