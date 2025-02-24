package middleware

import (
	"github.com/gin-gonic/gin"
	"mySimpleFprum/internal/configs"
	"mySimpleFprum/pkg/jwt"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	header := c.Request.Header
	token := header.Get("Authorization")

	if strings.TrimSpace(token) == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	config := configs.Get()
	id, username, err := jwt.ValidateToken(token, config.Service.SecretJWT)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Set("id", id)
	c.Set("username", username)
	c.Next()
}
