package posts

import (
	"context"
	"mySimpleFprum/internal/model/posts"
	"strings"
	"time"
)

func (s *Service) CreatePost(ctx context.Context, userID int64, userName string, req posts.CreatePostRequest) error {
	now := time.Now()
	hashString := strings.Join(req.PostHashtags, ";")
	postModel := posts.PostModel{
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: hashString,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    userName,
		UpdatedBy:    userName,
		UserID:       userID,
	}

	err := s.repo.CreatePosts(ctx, postModel)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) GetPostDetail(ctx context.Context, userId, postId int64) (posts.Post, error) {
	likeCount, err := s.repo.GetPostLikeCount(ctx, postId)
	if err != nil {
		return posts.Post{}, err
	}
	comments, err := s.repo.GetPostComment(ctx, postId)
	if err != nil {
		return posts.Post{}, err
	}

	postDetail, err := s.repo.GetPost(ctx, userId, postId)
	if err != nil {
		return posts.Post{}, err
	}

	return posts.Post{
		PostDetail: postDetail,
		Comments:   comments,
		LikeCount:  likeCount,
	}, nil
}
