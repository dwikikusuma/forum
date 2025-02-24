package posts

import (
	"github.com/gin-gonic/gin"
	"mySimpleFprum/internal/model/posts"
	"net/http"
	"strconv"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var req posts.CreatePostRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	userId := c.GetInt64("id")
	userName := c.GetString("username")
	err = h.postSvc.CreatePost(ctx, userId, userName, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok": true,
	})
}

func (h *Handler) GetPost(c *gin.Context) {
	userId := c.GetInt64("id")
	postIdStr := c.Param("post_id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	post, err := h.postSvc.GetPostDetail(ctx, userId, postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": post,
	})
}
