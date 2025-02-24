package posts

import (
	"context"
	"database/sql"
	"errors"
	"mySimpleFprum/internal/model/posts"
)

func (r *Repository) CreateComment(ctx context.Context, commentModel posts.CommentModel) error {
	query := `
		INSERT INTO comments (post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by)
		VALUES (?, ?, ?, ?, ?, ?, ?)
`
	_, err := r.db.ExecContext(ctx, query,
		commentModel.PostID, commentModel.UserID, commentModel.CommentContent,
		commentModel.CreatedAt, commentModel.UpdatedAt, commentModel.CreatedBy, commentModel.UpdatedBy,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetPostComment(ctx context.Context, postId int64) ([]posts.CommentResponse, error) {
	query := `SELECT user_id, comment_content, created_at FROM comments WHERE post_id = ?`
	rows, err := r.db.QueryContext(ctx, query, postId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []posts.CommentResponse{}, err
		}
	}

	comments := make([]posts.CommentResponse, 0)
	for rows.Next() {
		comment := posts.CommentResponse{}
		err = rows.Scan(&comment.UserID, &comment.CommentContent, &comment.CreatedAt)
		if err != nil {
			return []posts.CommentResponse{}, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
