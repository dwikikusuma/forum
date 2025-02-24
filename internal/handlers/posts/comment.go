package posts

import (
	"github.com/gin-gonic/gin"
	"mySimpleFprum/internal/model/posts"
	"net/http"
	"strconv"
)

func (h *Handler) CreateComment(c *gin.Context) {

	var reqModel posts.CreateCommentRequest
	err := c.ShouldBindJSON(&reqModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userId := c.GetInt64("id")
	userName := c.GetString("username")
	postIdStr := c.Param("post_id")

	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.postSvc.CreateComment(c.Request.Context(), reqModel, userId, postId, userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
