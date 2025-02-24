package memberships

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mySimpleFprum/internal/model/memberships"
	"time"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Username, req.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email already exist")
	}

	passString, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return errHash
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(passString),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}
	return s.membershipRepo.CreateUser(ctx, model)
}
