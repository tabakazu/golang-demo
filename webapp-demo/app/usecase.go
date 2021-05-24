package app

import (
	"github.com/tabakazu/webapi-app/app/input"
	"github.com/tabakazu/webapi-app/domain/model"
)

type RegisterUserAccountService interface {
	Execute(input.RegisterUserAccountParam) (*model.UserAccountEntity, error)
}
