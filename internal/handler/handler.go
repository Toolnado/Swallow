package handler

import (
	"github.com/Toolnado/SwalloW/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/", h.GetAllUsers)
			users.GET("/:id", h.GetUser)
			users.GET("/my_account", h.GetMyAccount)

			posts := users.Group("posts")
			{
				posts.GET("/", h.GetAllPosts)
			}

			userPosts := users.Group("/:id/posts")
			{
				userPosts.POST("/", h.CreatePost)
				userPosts.GET("/", h.GetAllPostsThisUser)
				userPosts.GET("/:post_id", h.GetPost)
			}
		}
	}

	return router
}
