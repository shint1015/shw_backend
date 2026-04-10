package usecase

import (
	"context"
	"errors"
	"shwgrpc/internal/family/domain"
	"shwgrpc/internal/family/port"
)

type FamilyUsecase struct {
	familyRepo port.FamilyRepository
}

type FamilyRoleUsecase struct {
	familyRoleRepo port.FamilyRoleRepository
}

func NewFamilyUsecase(
	familyRepo port.FamilyRepository,
) *FamilyUsecase {
	return &FamilyUsecase{
		familyRepo: familyRepo,
	}
}

func NewFamilyRoleUsecase(
	familyRoleRepo port.FamilyRoleRepository,
) *FamilyRoleUsecase {
	return &FamilyRoleUsecase{
		familyRoleRepo: familyRoleRepo,
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

func (u *FamilyUsecase) AddFamilyMember(ctx context.Context, input port.AddFamilyMemberInput) error {
    return errors.New("not implemented")
}

func (u *FamilyUsecase) AcceptInvitation(ctx context.Context) error {
    return errors.New("not implemented")
}


func (u *FamilyRoleUsecase) GetRole(ctx context.Context, familyID uint64) ([]domain.FamilyRole, error) {
	return u.familyRoleRepo.Get(ctx, familyID)
}

func (u *FamilyRoleUsecase) CreateRole(ctx context.Context, input port.CreateRoleInput) error {
	familyRole := domain.FamilyRole{
		Name: input.Name,
		FamilyID: input.FamilyID,
	}
	return u.familyRoleRepo.Create(ctx, familyRole)
}

func (u *FamilyRoleUsecase) UpdateRole(ctx context.Context, input port.UpdateRoleInput) error {
	familyRoleID := input.ID
	familyRole := domain.FamilyRole{
		ID: &familyRoleID,
		Name: input.Name,
		FamilyID: input.FamilyID,
	}
	return u.familyRoleRepo.Update(ctx, familyRole)
}

func (u *FamilyRoleUsecase) DeleteRole(ctx context.Context, familyRoleID uint64) error {
	return u.familyRoleRepo.Delete(ctx, familyRoleID)
}