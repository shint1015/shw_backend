package infra

import (
	"context"
	"shwgrpc/internal/family/domain"
	"shwgrpc/internal/family/port"
	"shwgrpc/model"
	sharedmapper "shwgrpc/internal/shared/mapper"
	"time"
)

type FamilyRepository struct {}

func NewFamilyRepository() port.FamilyRepository {
	return &FamilyRepository{}
}

func (r *FamilyRepository) Get(_ context.Context, familyID uint64) (domain.Family, error) {
	var res []model.Family
	if err := model.DB.
		Where(&model.Family{ID: uint(familyID)}).
		Find(&res).Error; err != nil {
		return domain.Family{}, err
	}
	return mapFamily(res), nil
}

func (r *FamilyRepository) Create(_ context.Context, family domain.Family) error {
	//todo get owner user id
	userId := uint(1)
	m := model.Family{
		Name: family.Name,
		OwnerUserID: userId,
	}
	return m.Create(nil)
}

func (r *FamilyRepository) Update(_ context.Context, family domain.Family) error {
	// TODO: check user id is for the family data
	m := model.Family{
		ID: family.ID,
		Name: family.Name,
	}
	return m.Update(nil)
}

func (r *FamilyRepository) Delete(_ context.Context, familyId uint64) error {
	// TODO: check user id is for the family data
	m := model.Family{ID: uint(familyId)}
	return m.Delete(nil)
}

type FamilyRoleRepository struct {}

func NewFamilyRoleRepository() port.FamilyRoleRepository {
	return &FamilyRoleRepository{}
}

func (r *FamilyRoleRepository) Get(_ context.Context, familyId uint64) (domain.FamilyRole, error) {
	var res []model.FamilyRole
	if err := model.DB.
		Where(&model.FamilyRole{FamilyID: uint(familyId)}).
		Find(&res).Error; err != nil {
		return domain.FamilyRole{}, err
	}
	return mapFamilyRole(res), nil
}

func (r *FamilyRoleRepository) Create(_ context.Context, familyRole domain.FamilyRole) error {
	m := model.FamilyRole{
		Name: familyRole.Name,
		FamilyID: uint(familyRole.FamilyID),
	}
	return m.Create(nil)
}

func (r *FamilyRoleRepository) Update(_ context.Context, familyRole domain.FamilyRole) error {
	m := model.FamilyRole{
		ID: familyRole.ID,
		Name: familyRole.Name,
		FamilyID: uint(familyRole.FamilyID),
	}
	return m.Update(nil)
}

func (r *FamilyRoleRepository) Delete(_ context.Context, familyRoleId uint64) error {
	m := model.FamilyRole{ID: uint(familyRoleId)}
	return m.Delete(nil)
}

func mapFamilyRole(f model.FamilyRole) domain.FamilyRole {
	return domain.FamilyRole{
		ID: uint64(f.ID),
		Name: f.Name,
		FamilyID: uint64(f.FamilyID),
	}
}

func mapFamilyRoleList(f []model.FamilyRole) []domain.FamilyRole {
	var familyRoles []domain.FamilyRole
	for _, familyRole := range f {
		familyRoles = append(familyRoles, mapFamilyRole(familyRole))
	}
	return familyRoles
}


func mapFamily(f model.Family) domain.Family {
	return domain.Family{
		ID: uint64(f.ID),
		Name: f.Name,
		Users: mapUserInfoList(f.Users),
	}
}

func mapUserInfoList(u model.Family.Users) domain.Family.Users {
	var users []domain.Family.Users
	for user := range u {
		userDomain := sharedmapper.MapUserInfoAs(u, func(id uint64, name string) domain.UserInfo {
			return domain.UserInfo{
				ID:   id,
				Name: name,
			}
		})
		users.append()
	}
}

