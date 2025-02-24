package memberships

import (
	"context"
	"mySimpleFprum/internal/configs"
	"mySimpleFprum/internal/model/memberships"
	"time"
)

type membershipRepository interface {
	GetUser(ctx context.Context, username string, email string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
	CreateRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshToken(ctx context.Context, userId int64, time time.Time) (*memberships.RefreshTokenModel, error)
}

type service struct {
	configs        *configs.Config
	membershipRepo membershipRepository
}

func NewService(config *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		configs:        config,
		membershipRepo: membershipRepo,
	}
}
