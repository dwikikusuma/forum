package posts

import (
	"context"
	"mySimpleFprum/internal/model/posts"
	"time"
)

func (s *Service) CreateComment(ctx context.Context, req posts.CreateCommentRequest, userId, postId int64, username string) error {

	err := s.repo.IsPostExists(ctx, postId)
	if err != nil {
		return err
	}

	now := time.Now()
	model := posts.CommentModel{
		PostID:         postId,
		UserID:         userId,
		CommentContent: req.CommentContent,
		CreatedBy:      username,
		UpdatedBy:      username,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	err = s.repo.CreateComment(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
