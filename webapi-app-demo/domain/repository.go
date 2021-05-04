package domain

import (
	"context"

	"github.com/tabakazu/webapi-app/domain/entity"
)

type UserRepository interface {
	FindByEmail(context.Context, string) (*entity.User, error)
}
