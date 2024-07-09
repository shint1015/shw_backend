package main

import (
	"connectrpc.com/connect"
	"context"
	shwgrpc "shwgrpc/pkg/grpc"
)

func (s *ShwServer) GetHousework(ctx context.Context, req *connect.Request[shwgrpc.HouseworkRequest]) (*connect.Response[shwgrpc.HouseworkResponse], error) {
	return houseworkController.GetHousework(ctx, req)
}

func (s *ShwServer) GetHouseworkDetail(ctx context.Context, req *connect.Request[shwgrpc.HouseworkDetailRequest]) (*connect.Response[shwgrpc.HouseworkDetailResponse], error) {
	return houseworkController.GetHouseworkDetail(ctx, req)
}

func (s *ShwServer) CreateHousework(ctx context.Context, req *connect.Request[shwgrpc.Housework]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.CreateHousework(ctx, req)
}

func (s *ShwServer) UpdateHousework(ctx context.Context, req *connect.Request[shwgrpc.Housework]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.UpdateHousework(ctx, req)
}

func (s *ShwServer) FinishHousework(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTargetRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.FinishHousework(ctx, req)
}

func (s *ShwServer) DeleteHousework(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTargetRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.DeleteHousework(ctx, req)
}

func (s *ShwServer) GetHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemoRequest]) (*connect.Response[shwgrpc.HouseworkMemoResponse], error) {
	return houseworkController.GetHouseworkMemo(ctx, req)
}

func (s *ShwServer) CreateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.CreateHouseworkMemo(ctx, req)
}

func (s *ShwServer) UpdateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.UpdateHouseworkMemo(ctx, req)
}

func (s *ShwServer) DeleteHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.DeleteHouseworkMemo(ctx, req)
}

func (s *ShwServer) GetHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplateRequest]) (*connect.Response[shwgrpc.HouseworkTemplateResponse], error) {
	return houseworkController.GetHouseworkTemplate(ctx, req)
}

func (s *ShwServer) GetHouseworkTemplates(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplatesRequest]) (*connect.Response[shwgrpc.HouseworkTemplatesResponse], error) {
	return houseworkController.GetHouseworkTemplates(ctx, req)
}

func (s *ShwServer) CreateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.CreateHouseworkTemplate(ctx, req)
}

func (s *ShwServer) UpdateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.UpdateHouseworkTemplate(ctx, req)
}

func (s *ShwServer) DeleteHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.DeleteHouseworkTemplate(ctx, req)
}

func (s *ShwServer) GetHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.HouseworkPointRequest]) (*connect.Response[shwgrpc.HouseworkPointResponse], error) {
	return houseworkController.GetHouseworkPoint(ctx, req)
}

func (s *ShwServer) GetHouseworkPointHistory(ctx context.Context, req *connect.Request[shwgrpc.HouseworkPointHistoryRequest]) (*connect.Response[shwgrpc.HouseworkPointHistoryResponse], error) {
	return houseworkController.GetHouseworkPointHistory(ctx, req)
}

func (s *ShwServer) GetFamilyPointList(ctx context.Context, req *connect.Request[shwgrpc.FamilyPointListRequest]) (*connect.Response[shwgrpc.FamilyPointList], error) {
	return pointController.GetFamilyPointList(ctx, req)
}

func (s *ShwServer) GetFamily(ctx context.Context, req *connect.Request[shwgrpc.FamilyRequest]) (*connect.Response[shwgrpc.FamilyResponse], error) {
	return familyController.GetFamily(ctx, req)
}

func (s *ShwServer) CreateFamily(ctx context.Context, req *connect.Request[shwgrpc.Family]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.CreateFamily(ctx, req)
}

func (s *ShwServer) UpdateFamily(ctx context.Context, req *connect.Request[shwgrpc.Family]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.UpdateFamily(ctx, req)
}

func (s *ShwServer) DeleteFamily(ctx context.Context, req *connect.Request[shwgrpc.Family]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.DeleteFamily(ctx, req)
}

func (s *ShwServer) AddFamilyMember(ctx context.Context, req *connect.Request[shwgrpc.AddFamilyMemberRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.AddFamilyMember(ctx, req)
}

func (s *ShwServer) UpdateRole(ctx context.Context, req *connect.Request[shwgrpc.UpdateRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return userController.UpdateRole(ctx, req)
}

func (s *ShwServer) GetFamilyRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRequest]) (*connect.Response[shwgrpc.FamilyRoleResponse], error) {
	return familyController.GetRole(ctx, req)
}
func (s *ShwServer) CreateFamilyRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRole]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.CreateRole(ctx, req)
}
func (s *ShwServer) UpdateFamilyRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRole]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.UpdateRole(ctx, req)
}
func (s *ShwServer) DeleteFamilyRole(ctx context.Context, req *connect.Request[shwgrpc.FamilyRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.DeleteRole(ctx, req)
}

func (s *ShwServer) GetBelongToUser(ctx context.Context, req *connect.Request[shwgrpc.GetBelongToUserRequest]) (*connect.Response[shwgrpc.GetBelongToUserResponse], error) {
	return familyController.GetBelongToUser(ctx, req)
}
