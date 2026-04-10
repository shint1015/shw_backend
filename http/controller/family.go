package controller

import (
	"context"
	"shwgrpc/internal/family/domain"
	"shwgrpc/internal/family/infra"
	"shwgrpc/internal/family/port"
	"shwgrpc/internal/family/usecase"
	shwgrpc "shwgrpc/pkg/grpc"

	"connectrpc.com/connect"
)

type FamilyController struct{}

var familyUsecase port.FamilyUsecase = usecase.NewFamilyUsecase(
	infra.NewFamilyRepository(),
)

var familyRoleUsecase port.FamilyRoleUsecase = usecase.NewFamilyRoleUsecase(
	infra.NewFamilyRoleRepository(),
)

func NewFamilyController() *FamilyController {
	return &FamilyController{}
}

func (c *FamilyController) GetFamily(ctx context.Context, req *connect.Request[shwgrpc.GetFamilyRequest]) (*connect.Response[shwgrpc.GetFamilyResponse], error) {
	family, err := familyUsecase.GetFamily(ctx, req.Msg.FamilyId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&shwgrpc.GetFamilyResponse{Family: mapFamilyDomainToGrpc(family)})
	return res, nil
}

func (c *FamilyController) CreateFamily(ctx context.Context, req *connect.Request[shwgrpc.CreateFamilyRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyUsecase.CreateFamily(ctx, mapCreateFamilyInput(req.Msg.Family)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) UpdateFamily(ctx context.Context, req *connect.Request[shwgrpc.UpdateFamilyRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyUsecase.UpdateFamily(ctx, mapUpdateFamilyInput(req.Msg.Family)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) DeleteFamily(ctx context.Context, req *connect.Request[shwgrpc.DeleteFamilyRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyUsecase.DeleteFamily(ctx, req.Msg.FamilyId); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) AddFamilyMember(ctx context.Context, req *connect.Request[shwgrpc.AddFamilyMemberRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyUsecase.AddFamilyMember(ctx, mapAddFamilyMemberInput(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) GetRole(ctx context.Context, req *connect.Request[shwgrpc.GetFamilyRoleRequest]) (*connect.Response[shwgrpc.GetFamilyRoleResponse], error) {
	roles, err := familyRoleUsecase.GetRole(ctx, req.Msg.FamilyId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.GetFamilyRoleResponse{Roles: mapFamilyRoleDomainToGrpc(roles)})
	return res, nil
}
func (c *FamilyController) CreateRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRole]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyRoleUsecase.CreateRole(ctx, mapCreateFamilyRoleInput(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}
func (c *FamilyController) UpdateRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRole]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyRoleUsecase.UpdateRole(ctx, mapUpdateFamilyRoleInput(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}
func (c *FamilyController) DeleteRole(ctx context.Context, req *connect.Request[shwgrpc.DeleteFamilyRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyRoleUsecase.DeleteRole(ctx, req.Msg.FamilyRoleId); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func mapFamilyDomainToGrpc(f domain.Family) *shwgrpc.Family {
	return &shwgrpc.Family{
		Id: uint64(*f.ID),
		Name: f.Name,
	}
}

func mapCreateFamilyInput(f *shwgrpc.Family) port.CreateFamilyInput {
	if f == nil {
		return port.CreateFamilyInput{}
	}
	return port.CreateFamilyInput{
		Name: f.Name,
	}
}

func mapUpdateFamilyInput(f *shwgrpc.Family) port.UpdateFamilyInput {
	if f == nil {
		return port.UpdateFamilyInput{}
	}
	return port.UpdateFamilyInput{
		ID:   f.Id,
		Name: f.Name,
	}
}

func mapAddFamilyMemberInput(f *shwgrpc.AddFamilyMemberRequest) port.AddFamilyMemberInput {
	if f == nil {
		return port.AddFamilyMemberInput{}
	}
	return port.AddFamilyMemberInput{
		Name: f.Name,
		Email: f.Email,
		FamilyID: f.FamilyId,
	}
}

func mapCreateFamilyRoleInput(f *shwgrpc.FamilyRole) port.CreateRoleInput {
	if f == nil {
		return port.CreateRoleInput{}
	}
	return port.CreateRoleInput{
		Name: f.Name,
		FamilyID: f.FamilyId,
	}
}

func mapUpdateFamilyRoleInput(f *shwgrpc.FamilyRole) port.UpdateRoleInput {
	if f == nil {
		return port.UpdateRoleInput{}
	}
	return port.UpdateRoleInput{
		ID: f.Id,
		Name: f.Name,
		FamilyID: f.FamilyId,
	}
}
func mapFamilyRoleDomainToGrpc(frList []domain.FamilyRole) []*shwgrpc.FamilyRole {
	if frList == nil {
		return []*shwgrpc.FamilyRole{}
	}
	var grpcList []*shwgrpc.FamilyRole
	for _, fr := range frList {
		grpcList = append(grpcList, &shwgrpc.FamilyRole{
			Id: uint64(*fr.ID),
			Name: fr.Name,
			FamilyId: fr.FamilyID,
		})
	}
	return grpcList
}