package domain

import (
	"github.com/tabakazu/webapi-app/domain/model"
)

type UserAccountRepository interface {
	Create(*model.UserAccountEntity) error
	FindByID(*model.UserAccountEntity, uint) error
	FindByEmail(*model.UserAccountEntity, string) error
}
