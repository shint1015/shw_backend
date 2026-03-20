package usecase

import (
	"context"
	"shwgrpc/internal/family/domain"
	"shwgrpc/internal/family/port"
)

type FamilyUsecase struct {
	familyRepo port.FamilyRepository
}

func NewFamilyUsecase(
	familyRepo port.FamilyRepository,
) *FamilyUsecase {
	return &FamilyUsecase{
		familyRepo: familyRepo,
	}
}

func (u *FamilyUsecase) GetFamily(ctx context.Context, familyID uint64) (domain.Family, error) {
	return u.familyRepo.Get(ctx, familyID)
}


func (u *FamilyUsecase) CreateFamily(ctx context.Context, input port.CreateFamilyInput) error {
	family := domain.Family{
		Name: input.Name,
	}
	return u.familyRepo.Create(ctx, family)
}

func (u *FamilyUsecase) UpdateFamily(ctx context.Context, input port.UpdateFamilyInput) error {
	familyID := input.ID
	family := domain.Family{
		ID:   &familyID,
		Name: input.Name,
	}
	return u.familyRepo.Update(ctx, family)
}

func (u *FamilyUsecase) DeleteFamily(ctx context.Context, familyID uint64) error {
	return u.familyRepo.Delete(ctx, familyID)
}