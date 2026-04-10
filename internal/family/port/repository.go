package port

import (
	"context"
	"shwgrpc/internal/family/domain"
)

type FamilyRepository interface {
	Create(ctx context.Context, family domain.Family) error
	Update(ctx context.Context, family domain.Family) error
	Delete(ctx context.Context, familyId uint64) error
	Get(ctx context.Context, familyId uint64) (domain.Family, error)
}

type FamilyRoleRepository interface {
	Create(ctx context.Context, familyRole domain.FamilyRole) error
	Update(ctx context.Context, familyRole domain.FamilyRole) error
	Delete(ctx context.Context, familyRoleId uint64) error
	Get(ctx context.Context, familyId uint64) ([]domain.FamilyRole, error)
}