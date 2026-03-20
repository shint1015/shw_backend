package controller

import (
	"context"
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

func NewFamilyController() *FamilyController {
	return &FamilyController{}
}

func (c *FamilyController) GetFamily(ctx context.Context, req *connect.Request[shwgrpc.GetFamilyRequest]) (*connect.Response[shwgrpc.GetFamilyResponse], error) {
	family, err := familyUsecase.GetFamily(ctx, req.Msg.FamilyId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.GetFamilyResponse{Family: family})
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
	if err := familyUsecase.UpdateFamily(mapUpdateFamilyInput(req.Msg.Family)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) DeleteFamily(ctx context.Context, req *connect.Request[shwgrpc.DeleteFamilyRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyUsecase.DeleteFamily(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) AddFamilyMember(ctx context.Context, req *connect.Request[shwgrpc.AddFamilyMemberRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyUsecase.AddFamilyMember(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) GetRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRequest]) (*connect.Response[shwgrpc.FamilyRoleResponse], error) {
	roles, err := familyRoleUsecase.GetRole(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.FamilyRoleResponse{FamilyRole: roles})
	return res, nil
}
func (c *FamilyController) CreateRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRole]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyRoleUsecase.CreateRole(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}
func (c *FamilyController) UpdateRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRole]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyRoleUsecase.UpdateRole(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}
func (c *FamilyController) DeleteRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyRoleUsecase.DeleteRole(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
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
