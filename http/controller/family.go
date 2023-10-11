package controller

import (
	"context"
	"shwgrpc/internal/service"
	shwgrpc "shwgrpc/pkg/grpc"
)

type FamilyController struct{}

var familyService = service.NewFamilyService()

func NewFamilyController() *FamilyController {
	return &FamilyController{}
}

func (c *FamilyController) GetFamily(ctx context.Context, req *shwgrpc.FamilyRequest) (*shwgrpc.FamilyResponse, error) {
	family, err := familyService.GetFamily(req)
	if err != nil {
		return nil, err
	}
	return &shwgrpc.FamilyResponse{Family: family}, nil
}

func (c *FamilyController) CreateFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.FamilyCommonResponse, error) {
	if err := familyService.CreateFamily(req); err != nil {
		return nil, err
	}
	return &shwgrpc.FamilyCommonResponse{Message: "success"}, nil
}

func (c *FamilyController) UpdateFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.FamilyCommonResponse, error) {
	if err := familyService.UpdateFamily(req); err != nil {
		return nil, err
	}
	return &shwgrpc.FamilyCommonResponse{Message: "success"}, nil
}

func (c *FamilyController) DeleteFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.FamilyCommonResponse, error) {
	if err := familyService.DeleteFamily(req); err != nil {
		return nil, err
	}
	return &shwgrpc.FamilyCommonResponse{Message: "success"}, nil
}

func (c *FamilyController) GetFamilyHouseworkPoints(ctx context.Context, req *shwgrpc.FamilyHouseworkPointRequest) (*shwgrpc.FamilyHouseworkPointResponse, error) {
	familyHouseworkPoints, err := familyService.GetFamilyHouseworkPoints(req)
	if err != nil {
		return nil, err
	}
	return &shwgrpc.FamilyHouseworkPointResponse{HouseworkPoint: familyHouseworkPoints}, nil
}

func (c *FamilyController) AddFamilyMember(ctx context.Context, req *shwgrpc.AddFamilyMemberRequest) (*shwgrpc.FamilyCommonResponse, error) {
	if err := familyService.AddFamilyMember(req); err != nil {
		return nil, err
	}
	return &shwgrpc.FamilyCommonResponse{Message: "success"}, nil
}
