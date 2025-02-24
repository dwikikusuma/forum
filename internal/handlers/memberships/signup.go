package memberships

import (
	"github.com/gin-gonic/gin"
	"mySimpleFprum/internal/model/memberships"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var signUpReq memberships.SignUpRequest
	if err := c.ShouldBindJSON(&signUpReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := h.membershipSvc.SignUp(ctx, signUpReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok": true,
	})
}
