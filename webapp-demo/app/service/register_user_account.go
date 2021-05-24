package service

import (
	"errors"

	"github.com/tabakazu/webapi-app/app"
	"github.com/tabakazu/webapi-app/app/input"
	"github.com/tabakazu/webapi-app/domain"
	"github.com/tabakazu/webapi-app/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type registerUserAccount struct {
	repo domain.UserAccountRepository
}

func NewRegisterUserAccount(repo domain.UserAccountRepository) app.RegisterUserAccountService {
	return &registerUserAccount{repo: repo}
}

func (s *registerUserAccount) Execute(p input.RegisterUserAccountParam) (*model.UserAccountEntity, error) {
	if p.Password != p.PasswordConfirmation {
		return nil, errors.New("password doesn't match password_confirmation")
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	e := &model.UserAccountEntity{
		FamilyName:     p.FamilyName,
		GivenName:      p.GivenName,
		Email:          p.Email,
		PasswordDigest: string(passHash),
	}

	if err := s.repo.Create(e); err != nil {
		return nil, err
	}
	return e, nil
}
