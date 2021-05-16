package service

import (
	"context"
	"errors"

	"github.com/tabakazu/webapi-app/domain"
	"github.com/tabakazu/webapi-app/domain/entity"
	"github.com/tabakazu/webapi-app/domain/value"
)

type userSignInService struct {
	repo domain.UserRepository
}

func (s userSignInService) Execute(ctx context.Context, email, pass string) (*entity.User, error) {
	usr, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return usr, err
	}

	ok := usr.IsCorrectPassword(value.Password(pass))
	if !ok {
		return usr, errors.New("Password is invalid")
	}

	return usr, nil
}
