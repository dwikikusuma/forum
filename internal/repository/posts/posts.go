package posts

import (
	"context"
	"database/sql"
	"errors"
	"mySimpleFprum/internal/model/posts"
	"strings"
)

func (r *Repository) CreatePosts(ctx context.Context, model posts.PostModel) error {
	query := `
		INSERT INTO posts (
		user_id, post_title, post_content, post_hashtags, 
		created_at, updated_at, created_by, updated_by ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(
		ctx, query,
		model.UserID, model.PostTitle, model.PostContent, model.PostHashtags,
		model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) IsPostExists(ctx context.Context, postId int64) error {
	query := `SELECT id from posts WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, postId)

	var rowId int64
	err := row.Scan(&rowId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("post doesnt exists")
		}
		return err
	}

	return nil
}

func (r *Repository) GetPost(ctx context.Context, userId, postId int64) (posts.PostResponse, error) {
	query := `
		SELECT 
			p.id, p.user_id, p.post_title, p.post_content, p.post_hashtags, 
			p.created_at, p.updated_at, p.created_by, p.updated_by, COALESCE(ua.is_liked, 0) as "is_like", COALESCE (ua.is_saved, 0) as "is_saved"
		FROM posts p LEFT JOIN user_activities ua  on p.id = ua.post_id AND ua.user_id = ? 
		WHERE p.id = ?;`

	var model posts.PostResponse
	var hashtag string
	row := r.db.QueryRowContext(ctx, query, userId, postId)
	err := row.Scan(
		&model.ID, &model.UserID, &model.PostTitle, &model.PostContent, &hashtag,
		&model.CreatedAt, &model.UpdatedAt, &model.CreatedBy, &model.UpdatedBy, &model.IsLiked, &model.IsSaved,
	)

	model.PostHashtags = strings.Split(hashtag, ";")
	if err != nil {
		return posts.PostResponse{}, err
	}

	return model, nil
}

func (r *Repository) GetPostLikeCount(ctx context.Context, postId int64) (int64, error) {
	query := `SELECT COUNT(id) as total_likes FROM user_activities WHERE post_id = ?`
	var likeCount int64

	row := r.db.QueryRowContext(ctx, query, postId)
	err := row.Scan(&likeCount)
	if err != nil {
		return 0, err
	}

	return likeCount, nil
}
