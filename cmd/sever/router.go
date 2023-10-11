package main

import (
	"context"
	"shwgrpc/http/controller"

	shwgrpc "shwgrpc/pkg/grpc"
)

var houseworkController = controller.NewHouseworkController()
var familyController = controller.NewFamilyController()
var userController = controller.NewUserController()

func (s *ShwServer) GetHousework(ctx context.Context, req *shwgrpc.HouseworkRequest) (*shwgrpc.HouseworkResponse, error) {
	return houseworkController.GetHousework(ctx, req)
}

func (s *ShwServer) GetHouseworkDetail(ctx context.Context, req *shwgrpc.HouseworkDetailRequest) (*shwgrpc.HouseworkDetailResponse, error) {
	return houseworkController.GetHouseworkDetail(ctx, req)
}

func (s *ShwServer) CreateHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.CreateHousework(ctx, req)
}

func (s *ShwServer) UpdateHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.UpdateHousework(ctx, req)
}

func (s *ShwServer) FinishHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.FinishHousework(ctx, req)
}

func (s *ShwServer) DeleteHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.DeleteHousework(ctx, req)
}

func (s *ShwServer) GetHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemoRequest) (*shwgrpc.HouseworkMemoResponse, error) {
	return houseworkController.GetHouseworkMemo(ctx, req)
}

func (s *ShwServer) CreateHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.CreateHouseworkMemo(ctx, req)
}

func (s *ShwServer) UpdateHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.UpdateHouseworkMemo(ctx, req)
}

func (s *ShwServer) DeleteHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.DeleteHouseworkMemo(ctx, req)
}

func (s *ShwServer) GetHouseworkTemplate(ctx context.Context, req *shwgrpc.HouseworkTemplateRequest) (*shwgrpc.HouseworkTemplateResponse, error) {
	return houseworkController.GetHouseworkTemplate(ctx, req)
}

func (s *ShwServer) CreateHouseworkTemplate(ctx context.Context, req *shwgrpc.HouseworkTemplate) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.CreateHouseworkTemplate(ctx, req)
}

func (s *ShwServer) UpdateHouseworkTemplate(ctx context.Context, req *shwgrpc.HouseworkTemplate) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.UpdateHouseworkTemplate(ctx, req)
}

func (s *ShwServer) RemoveHouseworkTemplate(ctx context.Context, req *shwgrpc.HouseworkTemplate) (*shwgrpc.HouseworkCommonResponse, error) {
	return houseworkController.RemoveHouseworkTemplate(ctx, req)
}

func (s *ShwServer) GetHouseworkPoint(ctx context.Context, req *shwgrpc.HouseworkPointRequest) (*shwgrpc.HouseworkPointResponse, error) {
	return houseworkController.GetHouseworkPoint(ctx, req)
}

func (s *ShwServer) GetHouseworkPointHistory(ctx context.Context, req *shwgrpc.HouseworkPointHistoryRequest) (*shwgrpc.HouseworkPointHistoryResponse, error) {
	return houseworkController.GetHouseworkPointHistory(ctx, req)
}

func (s *ShwServer) GetFamily(ctx context.Context, req *shwgrpc.FamilyRequest) (*shwgrpc.FamilyResponse, error) {
	return familyController.GetFamily(ctx, req)
}

func (s *ShwServer) CreateFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.FamilyCommonResponse, error) {
	return familyController.CreateFamily(ctx, req)
}

func (s *ShwServer) UpdateFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.FamilyCommonResponse, error) {
	return familyController.UpdateFamily(ctx, req)
}

func (s *ShwServer) DeleteFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.FamilyCommonResponse, error) {
	return familyController.DeleteFamily(ctx, req)
}

func (s *ShwServer) AddFamilyMember(ctx context.Context, req *shwgrpc.AddFamilyMemberRequest) (*shwgrpc.FamilyCommonResponse, error) {
	return familyController.AddFamilyMember(ctx, req)
}

func (s *ShwServer) UpdateRole(ctx context.Context, req *shwgrpc.UpdateRoleRequest) (*shwgrpc.UserCommonResponse, error) {
	return userController.UpdateRole(ctx, req)
}
