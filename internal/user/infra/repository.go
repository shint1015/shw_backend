package infra

import (
	"context"
	"shwgrpc/internal/user/port"
	"shwgrpc/model"
)

type UserRepository struct{}

func NewUserRepository() port.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) UpdateRole(_ context.Context, userID uint64, roleID uint64) error {
	role := uint(roleID)
	u := model.User{
		ID:     uint(userID),
		RoleID: &role,
	}
	return u.Update(nil)
}
