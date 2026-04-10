package controller

import (
	"context"
	"shwgrpc/internal/housework/domain"
	"shwgrpc/internal/housework/infra"
	"shwgrpc/internal/housework/port"
	"shwgrpc/internal/housework/usecase"
	shwgrpc "shwgrpc/pkg/grpc"
	"time"

	"connectrpc.com/connect"
)

type HouseworkController struct{}

var houseworkUsecase port.HouseworkUsecase = usecase.NewHouseworkUsecase(
	infra.NewHouseworkRepository(),
	infra.NewHouseworkMemoRepository(),
	infra.NewHouseworkTemplateRepository(),
	infra.NewUserRepository(),
	infra.NewHouseworkPointHistoryRepository(),
)

func NewHouseworkController() *HouseworkController {
	return &HouseworkController{}
}

func (c *HouseworkController) GetHousework(ctx context.Context, req *connect.Request[shwgrpc.GetHouseworkRequest]) (*connect.Response[shwgrpc.GetHouseworkResponse], error) {
	housework, memo, err := houseworkUsecase.GetHouseworkDetail(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.GetHouseworkResponse{
		Housework: mapHouseworkToGrpc(housework),
		Memos:      mapHouseworkMemoListToGrpc(memo),
	})
	return res, nil
}

func (c *HouseworkController) ListHouseworks(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworksRequest]) (*connect.Response[shwgrpc.ListHouseworksResponse], error) {
	houseworks, err := houseworkUsecase.ListHousework(ctx, req.Msg.FamilyId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.ListHouseworksResponse{
		Houseworks: mapHouseworkListToGrpc(houseworks),
	})
	return res, nil
}

func (c *HouseworkController) CreateHousework(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.CreateHousework(ctx, mapHouseworkCreateInputFromGrpc(req.Msg.Housework)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHousework(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.UpdateHousework(ctx, mapHouseworkUpdateInputFromGrpc(req.Msg.Housework)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) FinishHousework(ctx context.Context, req *connect.Request[shwgrpc.FinishHouseworkRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.FinishHousework(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHousework(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.DeleteHousework(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) ListHouseworkMemos(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworkMemosRequest]) (*connect.Response[shwgrpc.ListHouseworkMemosResponse], error) {
	memoList, err := houseworkUsecase.ListHouseworkMemos(ctx, req.Msg.HouseworkId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.ListHouseworkMemosResponse{Memos: mapHouseworkMemoListToGrpc(memoList)})
	return res, nil
}

func (c *HouseworkController) CreateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkMemoRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.CreateHouseworkMemo(ctx, mapHouseworkMemoCreateInputFromGrpc(req.Msg.Memo)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkMemoRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.UpdateHouseworkMemo(ctx, mapHouseworkMemoUpdateInputFromGrpc(req.Msg.Memo)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkMemoRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.DeleteHouseworkMemo(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.GetHouseworkTemplateRequest]) (*connect.Response[shwgrpc.GetHouseworkTemplateResponse], error) {
	houseworkTemplate, err := houseworkUsecase.GetHouseworkTemplate(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.GetHouseworkTemplateResponse{Template: mapHouseworkTemplateToGrpc(houseworkTemplate)})
	return res, nil
}

func (c *HouseworkController) ListHouseworkTemplates(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworkTemplatesRequest]) (*connect.Response[shwgrpc.ListHouseworkTemplatesResponse], error) {
	houseworkTemplates, err := houseworkUsecase.ListHouseworkTemplates(ctx, req.Msg.FamilyId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.ListHouseworkTemplatesResponse{Templates: mapHouseworkTemplateListToGrpc(houseworkTemplates)})
	return res, nil
}

func (c *HouseworkController) CreateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkTemplateRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.CreateHouseworkTemplate(ctx, mapHouseworkTemplateCreateInputFromGrpc(req.Msg.Template)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkTemplateRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.UpdateHouseworkTemplate(ctx, mapHouseworkTemplateUpdateInputFromGrpc(req.Msg.Template)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkTemplateRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.DeleteHouseworkTemplate(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.GetHouseworkPointRequest]) (*connect.Response[shwgrpc.GetHouseworkPointResponse], error) {
	houseworkPoint, err := houseworkUsecase.GetHouseworkPoint(ctx, req.Msg.UserId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.GetHouseworkPointResponse{Point: mapHouseworkPointToGrpc(houseworkPoint)})
	return res, nil
}

func (c *HouseworkController) ListHouseworkPointHistories(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworkPointHistoriesRequest]) (*connect.Response[shwgrpc.ListHouseworkPointHistoriesResponse], error) {
	houseworkPointHistories, err := houseworkUsecase.ListHouseworkPointHistories(ctx, req.Msg.UserId)
	if err != nil {
		return nil, err
	}
	res := connect.NnewResponse(&shwgrpc.ListHouseworkPointHistoriesResponse{Histories: mapHouseworkPointHistoryListToGrpc(houseworkPointHistories)})
	return res, nil
}

func (c *HouseworkController) CreateHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkPointRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.CreateHouseworkPoint(ctx, mapHouseworkPointFromGrpc(req.Msg.Point)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkPointRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.UpdateHouseworkPoint(ctx, mapHouseworkPointFromGrpc(req.Msg.Point)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkPointRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.DeleteHouseworkPoint(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkSchedule(ctx context.Context, req *connect.Request[shwgrpc.GetHouseworkScheduleRequest]) (*connect.Response[shwgrpc.GetHouseworkScheduleResponse], error) {
	houseworkSchedule, err := houseworkUsecase.GetHouseworkSchedule(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.GetHouseworkScheduleResponse{Schedule: mapHouseworkScheduleToGrpc(houseworkSchedule)})
	return res, nil
}

func (c *HouseworkController) ListHouseworkSchedules(ctx context.Context, req *connect.Request[shwgrpc.ListHouseworkSchedulesRequest]) (*connect.Response[shwgrpc.ListHouseworkSchedulesResponse], error) {
	houseworkSchedules, err := houseworkUsecase.ListHouseworkSchedules(ctx, req.Msg.FamilyId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.ListHouseworkSchedulesResponse{Schedules: mapHouseworkScheduleListToGrpc(houseworkSchedules)})
	return res, nil
}

func (c *HouseworkController) CreateHouseworkSchedule(ctx context.Context, req *connect.Request[shwgrpc.CreateHouseworkScheduleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.CreateHouseworkSchedule(ctx, mapHouseworkScheduleFromGrpc(req.Msg.Schedule)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHouseworkSchedule(ctx context.Context, req *connect.Request[shwgrpc.UpdateHouseworkScheduleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.UpdateHouseworkSchedule(ctx, mapHouseworkScheduleFromGrpc(req.Msg.Schedule)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHouseworkSchedule(ctx context.Context, req *connect.Request[shwgrpc.DeleteHouseworkScheduleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.DeleteHouseworkSchedule(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func mapHouseworkListToGrpc(list []domain.Housework) []*shwgrpc.Housework {
	res := make([]*shwgrpc.Housework, 0, len(list))
	for _, item := range list {
		res = append(res, mapHouseworkToGrpc(item))
	}
	return res
}

func mapHouseworkToGrpc(h domain.Housework) *shwgrpc.Housework {
	return &shwgrpc.Housework{
		Id:       h.ID,
		FamilyId: h.FamilyID,
		Title:    h.Title,
		Detail:   h.Detail,
		StatusId:   h.StatusID,
		WorkUser: &shwgrpc.UserInfo{
			Id:   h.WorkUser.ID,
			Name: h.WorkUser.Name,
		},
		StartedAt: h.StartedAt.Unix(),
		EndedAt:   h.EndedAt.Unix(),
	}
}

func mapHouseworkCreateInputFromGrpc(h *shwgrpc.Housework) port.CreateHouseworkInput {
	if h == nil {
		return port.CreateHouseworkInput{}
	}
	return port.CreateHouseworkInput{
		FamilyID:   h.FamilyId,
		Title:      h.Title,
		Detail:     h.Detail,
		StatusID:   h.StatusId,
		WorkUserID: h.GetWorkUser().GetId(),
		StartedAt:  time.Unix(h.StartedAt, 0),
		EndedAt:    time.Unix(h.EndedAt, 0),
	}
}

func mapHouseworkUpdateInputFromGrpc(h *shwgrpc.Housework) port.UpdateHouseworkInput {
	if h == nil {
		return port.UpdateHouseworkInput{}
	}
	return port.UpdateHouseworkInput{
		ID:         h.Id,
		Title:      h.Title,
		Detail:     h.Detail,
		WorkUserID: h.GetWorkUser().GetId(),
		StartedAt:  time.Unix(h.StartedAt, 0),
		EndedAt:    time.Unix(h.EndedAt, 0),
	}
}

func mapHouseworkMemoListToGrpc(list []domain.HouseworkMemo) []*shwgrpc.HouseworkMemo {
	res := make([]*shwgrpc.HouseworkMemo, 0, len(list))
	for _, item := range list {
		res = append(res, mapHouseworkMemoToGrpc(item))
	}
	return res
}

func mapHouseworkMemoToGrpc(m domain.HouseworkMemo) *shwgrpc.HouseworkMemo {
	return &shwgrpc.HouseworkMemo{
		Id:          m.ID,
		HouseworkId: m.HouseworkID,
		Message:     m.Message,
		SendFrom: &shwgrpc.UserInfo{
			Id:   m.SendFrom.ID,
			Name: m.SendFrom.Name,
		},
	}
}

func mapHouseworkMemoCreateInputFromGrpc(m *shwgrpc.HouseworkMemo) port.CreateHouseworkMemoInput {
	if m == nil {
		return port.CreateHouseworkMemoInput{}
	}
	return port.CreateHouseworkMemoInput{
		HouseworkID: m.HouseworkId,
		Message:     m.Message,
		SendFromID:  m.GetSendFrom().GetId(),
	}
}

func mapHouseworkMemoUpdateInputFromGrpc(m *shwgrpc.HouseworkMemo) port.UpdateHouseworkMemoInput {
	if m == nil {
		return port.UpdateHouseworkMemoInput{}
	}
	return port.UpdateHouseworkMemoInput{
		ID:      m.Id,
		Message: m.Message,
	}
}

func mapHouseworkTemplateToGrpc(t domain.HouseworkTemplate) *shwgrpc.HouseworkTemplate {
	return &shwgrpc.HouseworkTemplate{
		Id:        t.ID,
		FamilyId:  t.FamilyID,
		Title:     t.Title,
		Detail:    t.Detail,
		CreatedAt: t.CreatedAt.Unix(),
		UpdatedAt: t.UpdatedAt.Unix(),
	}
}

func mapHouseworkTemplateCreateInputFromGrpc(t *shwgrpc.HouseworkTemplate) port.CreateHouseworkTemplateInput {
	if t == nil {
		return port.CreateHouseworkTemplateInput{}
	}
	return port.CreateHouseworkTemplateInput{
		FamilyID: t.FamilyId,
		Title:    t.Title,
		Detail:   t.Detail,
	}
}

func mapHouseworkTemplateUpdateInputFromGrpc(t *shwgrpc.HouseworkTemplate) port.UpdateHouseworkTemplateInput {
	if t == nil {
		return port.UpdateHouseworkTemplateInput{}
	}
	return port.UpdateHouseworkTemplateInput{
		ID:     t.Id,
		Title:  t.Title,
		Detail: t.Detail,
	}
}

func mapHouseworkTemplateListToGrpc(list []domain.HouseworkTemplate) []*shwgrpc.HouseworkTemplate {
	res := make([]*shwgrpc.HouseworkTemplate, 0, len(list))
	for _, item := range list {
		res = append(res, mapHouseworkTemplateToGrpc(item))
	}
	return res
}

func mapHouseworkPointToGrpc(point *domain.HouseworkPoint) *shwgrpc.HouseworkPoint {
	return &shwgrpc.HouseworkPoint{
		Point: point.Point,
		User: &shwgrpc.UserInfo{
			Id:   point.User.ID,
			Name: point.User.Name,
		},
		CreatedAt: point.CreatedAt.Unix(),
		UpdatedAt: point.UpdatedAt.Unix(),
	}
}

func mapHouseworkPointFromGrpc(p *shwgrpc.HouseworkPoint) *domain.HouseworkPoint {
	return &domain.HouseworkPoint{
		Point: p.Point,
		User: domain.UserInfo{
			ID:   p.User.Id,
			Name: p.User.Name,
		},
		CreatedAt: time.Unix(p.CreatedAt, 0),
		UpdatedAt: time.Unix(p.UpdatedAt, 0),
	}
}

func mapHouseworkPointHistoryListToGrpc(list []domain.HouseworkPointHistory) []*shwgrpc.HouseworkPointHistory {
	res := make([]*shwgrpc.HouseworkPointHistory, 0, len(list))
	for _, item := range list {
		res = append(res, &shwgrpc.HouseworkPointHistory{
			Id:        item.ID,
			Detail:    item.Detail,
			Point:     item.Point,
			CreatedAt: item.CreatedAt.Unix(),
		})
	}
	return res
}

