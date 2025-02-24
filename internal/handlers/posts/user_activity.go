package posts

import (
	"github.com/gin-gonic/gin"
	"log"
	"mySimpleFprum/internal/model/posts"
	"net/http"
	"strconv"
)

func (h *Handler) SetUserActivity(c *gin.Context) {
	userId := c.GetInt64("id")
	userName := c.GetString("username")

	postIdStr := c.Param("post_id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		log.Print(err)
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid post id",
		})
		return
	}

	var req posts.UserActivityRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx := c.Request.Context()

	err = h.postSvc.SetUserActivity(ctx, userId, postId, userName, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
