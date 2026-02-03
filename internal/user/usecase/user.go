package usecase

import (
	"context"
	"shwgrpc/internal/user/port"
)

type UserUsecase struct {
	userRepo port.UserRepository
}

func NewUserUsecase(userRepo port.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) UpdateRole(ctx context.Context, userID uint64, roleID uint64) error {
	return u.userRepo.UpdateRole(ctx, userID, roleID)
}
