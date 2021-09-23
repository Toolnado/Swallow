package handler

import (
	"net/http"

	"github.com/Toolnado/SwalloW/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var u model.User

	if err := c.BindJSON(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := h.service.Authentication.CreateUser(&u)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type transportStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var input transportStruct

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Authentication.GenerateToken(input.Username, input.Password)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
