package posts

import (
	"context"
	"github.com/gin-gonic/gin"
	"mySimpleFprum/internal/middleware"
	"mySimpleFprum/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, userName string, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, req posts.CreateCommentRequest, userId, postId int64, username string) error
	SetUserActivity(ctx context.Context, userId, postID int64, userName string, req posts.UserActivityRequest) error
	GetPostDetail(ctx context.Context, userId, postId int64) (posts.Post, error)
}

type Handler struct {
	engine  *gin.Engine
	postSvc postService
}

func NewHandler(engine *gin.Engine, postService postService) *Handler {
	return &Handler{
		engine:  engine,
		postSvc: postService,
	}
}

func (h *Handler) RegisterRoutes() {
	group := h.engine.Group("/posts")
	group.Use(middleware.AuthMiddleware)
	group.POST("/create", h.CreatePost)
	group.POST("/comment/:post_id", h.CreateComment)
	group.POST("/user_activity/:post_id", h.SetUserActivity)
	group.GET("/post-detail/:post_id", h.GetPost)
}
