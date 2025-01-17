package handlers

import (
	"errors"

	"backend.app/internal/contexts"
	"backend.app/internal/models"
	"github.com/gin-gonic/gin"
)

func GetSessionUser(c *gin.Context) (*models.User, error) {
	u, ok := c.Get(contexts.ContextUser)
	if !ok || u == nil {
		return nil, errors.New("no user found in context")
	}

	user, ok := u.(*models.User)
	if !ok {
		return nil, errors.New("unable to cast to session user")
	}

	return user, nil
}
