package main

import (
	"context"
	shwgrpc "shwgrpc/pkg/grpc"

	"connectrpc.com/connect"
)

func (s *ShwServer) GetHousework(ctx context.Context, req *connect.Request[shwgrpc.GetHouseworkRequest]) (*connect.Response[shwgrpc.GetHouseworkResponse], error) {
	return houseworkController.GetHousework(ctx, req)
}

func (s *ShwServer) ListHouseworks(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworksRequest]) (*connect.Response[shwgrpc.ListHouseworksResponse], error) {
	return houseworkController.ListHouseworks(ctx, req)
}

func (s *ShwServer) CreateHousework(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.CreateHousework(ctx, req)
}

func (s *ShwServer) UpdateHousework(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.UpdateHousework(ctx, req)
}

func (s *ShwServer) FinishHousework(ctx context.Context, req *connect.Request[shwgrpc.FinishHouseworkRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.FinishHousework(ctx, req)
}

func (s *ShwServer) DeleteHousework(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.DeleteHousework(ctx, req)
}

func (s *ShwServer) ListHouseworkMemos(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworkMemosRequest]) (*connect.Response[shwgrpc.ListHouseworkMemosResponse], error) {
	return houseworkController.ListHouseworkMemos(ctx, req)
}

func (s *ShwServer) CreateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkMemoRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.CreateHouseworkMemo(ctx, req)
}

func (s *ShwServer) UpdateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkMemoRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.UpdateHouseworkMemo(ctx, req)
}

func (s *ShwServer) DeleteHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkMemoRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.DeleteHouseworkMemo(ctx, req)
}

func (s *ShwServer) GetHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.GetHouseworkTemplateRequest]) (*connect.Response[shwgrpc.GetHouseworkTemplateResponse], error) {
	return houseworkController.GetHouseworkTemplate(ctx, req)
}

func (s *ShwServer) ListHouseworkTemplates(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworkTemplatesRequest]) (*connect.Response[shwgrpc.ListHouseworkTemplatesResponse], error) {
	return houseworkController.ListHouseworkTemplates(ctx, req)
}

func (s *ShwServer) CreateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkTemplateRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.CreateHouseworkTemplate(ctx, req)
}

func (s *ShwServer) UpdateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkTemplateRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.UpdateHouseworkTemplate(ctx, req)
}

func (s *ShwServer) DeleteHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkTemplateRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.DeleteHouseworkTemplate(ctx, req)
}

func (s *ShwServer) GetHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.GetHouseworkPointRequest]) (*connect.Response[shwgrpc.GetHouseworkPointResponse], error) {
	return houseworkController.GetHouseworkPoint(ctx, req)
}

func (s *ShwServer) ListHouseworkPointHistories(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworkPointHistoriesRequest]) (*connect.Response[shwgrpc.ListHouseworkPointHistoriesResponse], error) {
	return houseworkController.ListHouseworkPointHistories(ctx, req)
}

func (s *ShwServer) CreateHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkPointRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.CreateHouseworkPoint(ctx, req)
}

func (s *ShwServer) UpdateHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkPointRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.UpdateHouseworkPoint(ctx, req)
}

func (s *ShwServer) DeleteHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkPointRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.DeleteHouseworkPoint(ctx, req)
}


func (s *ShwServer) GetHouseworkSchedule(ctx context.Context, req *connect.Request[shwgrpc.GetHouseworkScheduleRequest]) (*connect.Response[shwgrpc.GetHouseworkScheduleResponse], error) {
	return houseworkController.GetHouseworkSchedule(ctx, req)
}

func (s *ShwServer) ListHouseworkSchedules(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworkSchedulesRequest]) (*connect.Response[shwgrpc.ListHouseworkSchedulesResponse], error) {
	return houseworkController.ListHouseworkSchedules(ctx, req)
}

func (s *ShwServer) CreateHouseworkSchedule(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkScheduleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.CreateHouseworkSchedule(ctx, req)
}

func (s *ShwServer) UpdateHouseworkSchedule(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkScheduleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.UpdateHouseworkSchedule(ctx, req)
}

func (s *ShwServer) DeleteHouseworkSchedule(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkScheduleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return houseworkController.DeleteHouseworkSchedule(ctx, req)
}

func (s *ShwServer) GetFamily(ctx context.Context, req *connect.Request[shwgrpc.GetFamilyRequest]) (*connect.Response[shwgrpc.GetFamilyResponse], error) {
	return familyController.GetFamily(ctx, req)
}

func (s *ShwServer) CreateFamily(ctx context.Context, req *connect.Request[shwgrpc.CreateFamilyRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.CreateFamily(ctx, req)
}

func (s *ShwServer) UpdateFamily(ctx context.Context, req *connect.Request[shwgrpc.UpdateFamilyRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.UpdateFamily(ctx, req)
}

func (s *ShwServer) DeleteFamily(ctx context.Context, req *connect.Request[shwgrpc.DeleteFamilyRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.DeleteFamily(ctx, req)
}

func (s *ShwServer) AddFamilyMember(ctx context.Context, req *connect.Request[shwgrpc.AddFamilyMemberRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.AddFamilyMember(ctx, req)
}

func (s *ShwServer) ListFamilyRoles(ctx context.Context, req *connect.Request[shwgrpc.ListFamilyRolesRequest]) (*connect.Response[shwgrpc.ListFamilyRolesResponse], error) {
	return familyController.ListFamilyRoles(ctx, req)
}

func (s *ShwServer) GetFamilyRole(ctx context.Context, req *connect.Request[shwgrpc.GetFamilyRoleRequest]) (*connect.Response[shwgrpc.GetFamilyRoleResponse], error) {
	return familyController.GetRole(ctx, req)
}
func (s *ShwServer) CreateFamilyRole(ctx context.Context, req *connect.Request[shwgrpc.CreateFamilyRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.CreateRole(ctx, req)
}
func (s *ShwServer) UpdateFamilyRole(ctx context.Context, req *connect.Request[shwgrpc.UpdateFamilyRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.UpdateRole(ctx, req)
}
func (s *ShwServer) DeleteFamilyRole(ctx context.Context, req *connect.Request[shwgrpc.DeleteFamilyRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.DeleteRole(ctx, req)
}

func (s *ShwServer) ListFamilyUsers(ctx context.Context, req *connect.Request[shwgrpc.ListFamilyUsersRequest]) (*connect.Response[shwgrpc.ListFamilyUsersResponse], error) {
	return familyController.GetBelongToUser(ctx, req)
}

func (s *ShwServer) UpdateUserRole(ctx context.Context, req *connect.Request[shwgrpc.UpdateUserRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	return familyController.GetBelongToUser(ctx, req)
}
