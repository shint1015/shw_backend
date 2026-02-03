package port

import "context"

type UserUsecase interface {
	UpdateRole(ctx context.Context, userID uint64, roleID uint64) error
}
