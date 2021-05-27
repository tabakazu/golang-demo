package service

import (
	"github.com/tabakazu/webapi-app/app"
	"github.com/tabakazu/webapi-app/app/input"
	"github.com/tabakazu/webapi-app/domain"
	"github.com/tabakazu/webapi-app/domain/model"
)

type showUserAccount struct {
	repo domain.UserAccountRepository
}

func NewShowUserAccount(repo domain.UserAccountRepository) app.ShowUserAccountService {
	return &showUserAccount{repo: repo}
}

func (s *showUserAccount) Execute(p input.ShowUserAccountParam) (*model.UserAccountEntity, error) {
	var e model.UserAccountEntity
	if err := s.repo.FindByID(&e, p.ID); err != nil {
		return nil, err
	}

	return &e, nil
}
