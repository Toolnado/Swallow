package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user_id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("emth auth header").Error())
		return
	}

	headersParts := strings.Split(header, " ")

	if len(headersParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("invalid auth header").Error())
		return
	}

	userID, err := h.service.Authentication.ParseToken(headersParts[1])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		c.AbortWithError(http.StatusInternalServerError, errors.New("user id not found"))

		return 0, errors.New("user id not found")
	}

	intId, ok := id.(int)

	if !ok {
		c.AbortWithError(http.StatusInternalServerError, errors.New("user id is of invalid type"))

		return 0, errors.New("user id is of invalid type")
	}

	return intId, nil
}
