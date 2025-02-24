package posts

import (
	"context"
	"database/sql"
	"errors"
	"mySimpleFprum/internal/model/posts"
)

func (r *Repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `INSERT INTO user_activities (
                             post_id, user_id, is_liked, is_saved,  
                             created_at, updated_at, created_by, updated_by
                             ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.IsLiked, model.IsSaved, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `UPDATE user_activities SET is_liked = ?, is_saved = ?, updated_at = ?, updated_by = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, model.IsLiked, model.IsSaved, model.UpdatedAt, model.UpdatedBy, model.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserActivity(ctx context.Context, postId, userId int64) (*posts.UserActivityModel, error) {
	query := `SELECT id, post_id, user_id, is_liked, is_saved, created_at, created_by, updated_at, updated_by FROM user_activities WHERE post_id = ? AND user_id = ?`
	var model posts.UserActivityModel

	row := r.db.QueryRowContext(ctx, query, postId, userId)
	err := row.Scan(
		&model.ID, &model.PostID, &model.UserID, &model.IsLiked, &model.IsSaved,
		&model.CreatedAt, &model.CreatedBy, &model.UpdatedAt, &model.UpdatedBy,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &model, nil
}
