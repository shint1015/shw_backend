package infra

import (
	"context"
	"shwgrpc/internal/family/domain"
	"shwgrpc/internal/family/port"
	houseworkdomain "shwgrpc/internal/housework/domain"
	sharedmapper "shwgrpc/internal/shared/mapper"
	"shwgrpc/model"
)

type FamilyRepository struct {}

func NewFamilyRepository() port.FamilyRepository {
	return &FamilyRepository{}
}

func (r *FamilyRepository) Get(_ context.Context, familyID uint64) (domain.Family, error) {
	var res model.Family
	if err := model.DB.
		Where(&model.Family{ID: uint(familyID)}).
		First(&res).Error; err != nil {
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
		ID: uint(*family.ID),
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

func (r *FamilyRoleRepository) Get(_ context.Context, familyId uint64) ([]domain.FamilyRole, error) {
	var res []model.FamilyRole
	if err := model.DB.
		Where(&model.FamilyRole{FamilyID: uint(familyId)}).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return mapFamilyRoleList(res), nil
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
		ID: uint(*familyRole.ID),
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
	id := uint64(f.ID)
	return domain.FamilyRole{
		ID: &id,
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
	id := uint64(f.ID)
	return domain.Family{
		ID: &id,
		Name: f.Name,
		Users: mapUserInfoList(f),
	}
}

func mapUserInfoList(u model.Family) []houseworkdomain.UserInfo {
	var users []houseworkdomain.UserInfo
	if u.Users == nil {
		return users
	}
	for _, user := range *u.Users {
		userDomain := sharedmapper.MapUserInfoAs(user, func(id uint64, name string) houseworkdomain.UserInfo {
			return houseworkdomain.UserInfo{
				ID:   id,
				Name: name,
			}
		})
		users = append(users, userDomain)
	}
	return users
}

