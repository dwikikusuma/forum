package posts

import (
	"context"
	"mySimpleFprum/internal/model/posts"
)

type postRepo interface {
	CreatePosts(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, commentModel posts.CommentModel) error
	IsPostExists(ctx context.Context, postId int64) error

	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	GetUserActivity(ctx context.Context, postId, userId int64) (*posts.UserActivityModel, error)

	GetPost(ctx context.Context, userId, postId int64) (posts.PostResponse, error)
	GetPostComment(ctx context.Context, postId int64) ([]posts.CommentResponse, error)
	GetPostLikeCount(ctx context.Context, postId int64) (int64, error)
}

type Service struct {
	repo postRepo
}

func NewService(repo postRepo) *Service {
	return &Service{
		repo: repo,
	}
}
