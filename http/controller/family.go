package controller

import (
	"connectrpc.com/connect"
	"context"
	"shwgrpc/internal/service"
	shwgrpc "shwgrpc/pkg/grpc"
)

type FamilyController struct{}

var familyService = service.NewFamilyService()

func NewFamilyController() *FamilyController {
	return &FamilyController{}
}

func (c *FamilyController) GetFamily(ctx context.Context, req *connect.Request[shwgrpc.FamilyRequest]) (*connect.Response[shwgrpc.FamilyResponse], error) {
	family, err := familyService.GetFamily(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.FamilyResponse{Family: family})
	return res, nil
}

func (c *FamilyController) CreateFamily(ctx context.Context, req *connect.Request[shwgrpc.Family]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyService.CreateFamily(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) UpdateFamily(ctx context.Context, req *connect.Request[shwgrpc.Family]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyService.UpdateFamily(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) DeleteFamily(ctx context.Context, req *connect.Request[shwgrpc.Family]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyService.DeleteFamily(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *FamilyController) GetFamilyHouseworkPoints(ctx context.Context, req *connect.Request[shwgrpc.FamilyHouseworkPointRequest]) (*connect.Response[shwgrpc.FamilyHouseworkPointResponse], error) {
	familyHouseworkPoints, err := familyService.GetFamilyHouseworkPoints(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.FamilyHouseworkPointResponse{HouseworkPoint: familyHouseworkPoints})
	return res, nil
}

func (c *FamilyController) AddFamilyMember(ctx context.Context, req *connect.Request[shwgrpc.AddFamilyMemberRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := familyService.AddFamilyMember(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}
