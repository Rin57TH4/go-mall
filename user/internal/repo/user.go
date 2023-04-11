package repo

import (
	"context"
	"user/internal/model"
)

type UserRepo interface {
	Save(ctx context.Context, data *model.User) error
}
