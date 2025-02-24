package posts

import (
	"context"
	"errors"
	"mySimpleFprum/internal/model/posts"
	"time"
)

func (s *Service) SetUserActivity(ctx context.Context, userId, postID int64, userName string, req posts.UserActivityRequest) error {
	postExistsErr := s.repo.IsPostExists(ctx, postID)
	if postExistsErr != nil {
		return errors.New("invalid post id")
	}

	userActivity, err := s.repo.GetUserActivity(ctx, postID, userId)
	if err != nil {
		return err
	}

	now := time.Now()
	model := posts.UserActivityModel{
		UserID:    userId,
		PostID:    postID,
		IsLiked:   req.IsLiked,
		IsSaved:   req.IsSaved,
		UpdatedAt: now,
		UpdatedBy: userName,
	}

	if userActivity == nil {
		model.CreatedAt = now
		model.CreatedBy = userName
		err = s.repo.CreateUserActivity(ctx, model)
		if err != nil {
			return err
		}
	} else {
		model.ID = userActivity.ID
		err = s.repo.UpdateUserActivity(ctx, model)
		if err != nil {
			return err
		}
	}

	return nil
}
