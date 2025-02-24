package memberships

import (
	"github.com/gin-gonic/gin"
	membershipsModel "mySimpleFprum/internal/model/memberships"
	"net/http"
)

func (h *Handler) Login(c *gin.Context) {
	var request membershipsModel.LoginRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	token, refreshToken, err := h.membershipSvc.Login(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := membershipsModel.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
