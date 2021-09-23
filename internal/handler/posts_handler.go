package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Toolnado/SwalloW/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPostsThisUser(c *gin.Context) {
	userID, err := getUserID(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	posts, err := h.service.Posts.GetAllPostsThisUser(id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[int][]model.Post{userID: posts})
}

func (h *Handler) GetAllPosts(c *gin.Context) {
	userID, err := getUserID(c)
	log.Print(userID)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	posts, err := h.service.Posts.GetAllPosts()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[int][]model.Post{userID: posts})
}

func (h *Handler) CreatePost(c *gin.Context) {
	userID, err := getUserID(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var p model.Post
	if err := c.BindJSON(&p); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p.User = userID

	id, err := h.service.Posts.CreatePost(&p)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) GetPost(c *gin.Context) {
	userID, err := getUserID(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	id, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	post, err := h.service.Posts.GetPost(id)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, map[int]model.Post{userID: post})
}
