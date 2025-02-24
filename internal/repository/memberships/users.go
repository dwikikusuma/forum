package memberships

import (
	"context"
	"database/sql"
	"errors"
	"mySimpleFprum/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, username string, email string) (*memberships.UserModel, error) {

	query := `
		SELECT 
		    id, 
		    email, 
		    password, 
		    username, 
		    created_at, 
		    updated_at, 
		    created_by, 
		    updated_by
		FROM users
		WHERE username = ? or email = ?
	`

	row := r.db.QueryRowContext(ctx, query, username, email)

	var response memberships.UserModel
	err := row.Scan(
		&response.ID,
		&response.Email,
		&response.Password,
		&response.Username,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.CreatedBy,
		&response.UpdatedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil

}
