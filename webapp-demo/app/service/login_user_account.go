package service

import (
	"github.com/tabakazu/webapi-app/app"
	"github.com/tabakazu/webapi-app/app/input"
	"github.com/tabakazu/webapi-app/domain"
	"github.com/tabakazu/webapi-app/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type loginUserAccount struct {
	repo domain.UserAccountRepository
}

func NewLoginUserAccount(repo domain.UserAccountRepository) app.LoginUserAccountService {
	return &loginUserAccount{repo: repo}
}

func (s *loginUserAccount) Execute(p input.LoginUserAccountParam) (*model.UserAccountEntity, error) {
	var e model.UserAccountEntity
	if err := s.repo.FindByEmail(&e, p.Email); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(e.PasswordDigest), []byte(p.Password)); err != nil {
		return nil, err
	}
	return &e, nil
}
