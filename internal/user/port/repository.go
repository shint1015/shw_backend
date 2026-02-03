package port

import "context"

type UserRepository interface {
	UpdateRole(ctx context.Context, userID uint64, roleID uint64) error
}
