package handler

import (
	"net/http"
	"strconv"

	"github.com/Toolnado/SwalloW/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
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

	user, err := h.service.Users.GetUserByID(id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[int]model.User{userID: user})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	userID, err := getUserID(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	users, err := h.service.Users.GetAllUsers()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[int][]model.User{userID: users})
}

func (h *Handler) GetMyAccount(c *gin.Context) {
	myID, err := getUserID(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user, err := h.service.Users.GetMyAccount(myID)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, map[int]model.User{myID: user})
}
