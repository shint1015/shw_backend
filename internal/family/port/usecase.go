package port

import (
	"context"
	"shwgrpc/internal/family/domain"
)

type FamilyUsecase interface {
	GetFamily(ctx context.Context, familyID uint64) (domain.Family, error)
	CreateFamily(ctx context.Context, input CreateFamilyInput) error
	UpdateFamily(ctx context.Context, input UpdateFamilyInput) error
	DeleteFamily(ctx context.Context, familyID uint64) error
	AddFamilyMember(ctx context.Context, )
}

type FamilyRoleUsecase interface {
	GetRole(ctx context.Context, ) (error)
	CreateRole(ctx context.Context, ) error
	UpdateRole(ctx context.Context, ) error
	DeleteRole(ctx context.Context, roleID uint64) error
}