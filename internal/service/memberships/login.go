package memberships

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mySimpleFprum/internal/model/memberships"
	"mySimpleFprum/pkg/jwt"
	internalToken "mySimpleFprum/pkg/token"
	"time"
)

func (s *service) Login(ctx context.Context, request memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, request.Username, request.Email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	token, err := jwt.GenerateToken(user.ID, user.Username, s.configs.Service.SecretJWT)
	if err != nil {
		return "", "", err
	}

	var refreshToken string

	existToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		return "", "", err
	}

	if existToken == nil {
		now := time.Now()
		newRefreshToken, err := internalToken.GenerateRefreshToken()
		if err != nil {
			return "", "", err
		}
		refreshTokenModel := memberships.RefreshTokenModel{
			UserID:       user.ID,
			RefreshToken: newRefreshToken,
			ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
			CreatedAt:    now,
			UpdatedAt:    now,
			CreatedBy:    user.Username,
			UpdatedBy:    user.Username,
		}
		err = s.membershipRepo.CreateRefreshToken(ctx, refreshTokenModel)
		if err != nil {
			return "", "", err
		}
		refreshToken = newRefreshToken
	} else {
		refreshToken = existToken.RefreshToken
	}

	return token, refreshToken, nil
}
