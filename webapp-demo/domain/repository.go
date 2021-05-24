package domain

import (
	"github.com/tabakazu/webapi-app/domain/model"
)

type UserAccountRepository interface {
	Create(*model.UserAccountEntity) error
}
