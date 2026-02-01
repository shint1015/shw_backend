package controller

import (
	"connectrpc.com/connect"
	"context"
	"shwgrpc/internal/housework/domain"
	"shwgrpc/internal/housework/infra"
	"shwgrpc/internal/housework/port"
	"shwgrpc/internal/housework/usecase"
	shwgrpc "shwgrpc/pkg/grpc"
	"time"
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

func (c *HouseworkController) GetHousework(ctx context.Context, req *connect.Request[shwgrpc.HouseworkRequest]) (*connect.Response[shwgrpc.HouseworkResponse], error) {
	housework, err := houseworkUsecase.ListHousework(ctx, req.Msg.FamilyId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkResponse{Housework: mapHouseworkListToGrpc(housework)})
	return res, nil
}

func (c *HouseworkController) GetHouseworkDetail(ctx context.Context, req *connect.Request[shwgrpc.HouseworkDetailRequest]) (*connect.Response[shwgrpc.HouseworkDetailResponse], error) {
	housework, houseworkMemos, err := houseworkUsecase.GetHouseworkDetail(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkDetailResponse{
		Housework: mapHouseworkToGrpc(housework),
		Memo:      mapHouseworkMemoListToGrpc(houseworkMemos),
	})
	return res, nil
}

func (c *HouseworkController) CreateHousework(ctx context.Context, req *connect.Request[shwgrpc.Housework]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.CreateHousework(ctx, mapHouseworkFromGrpc(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHousework(ctx context.Context, req *connect.Request[shwgrpc.Housework]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.UpdateHousework(ctx, mapHouseworkFromGrpc(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) FinishHousework(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTargetRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.FinishHousework(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHousework(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTargetRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.DeleteHousework(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemoRequest]) (*connect.Response[shwgrpc.HouseworkMemoResponse], error) {
	memoList, err := houseworkUsecase.ListHouseworkMemo(ctx, req.Msg.HouseworkId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkMemoResponse{Memo: mapHouseworkMemoListToGrpc(memoList)})
	return res, nil
}

func (c *HouseworkController) CreateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.CreateHouseworkMemo(ctx, mapHouseworkMemoFromGrpc(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.UpdateHouseworkMemo(ctx, mapHouseworkMemoFromGrpc(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.DeleteHouseworkMemo(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplateRequest]) (*connect.Response[shwgrpc.HouseworkTemplateResponse], error) {
	houseworkTemplate, err := houseworkUsecase.GetHouseworkTemplate(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkTemplateResponse{Template: mapHouseworkTemplateToGrpc(houseworkTemplate)})
	return res, nil
}

func (c *HouseworkController) GetHouseworkTemplates(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplatesRequest]) (*connect.Response[shwgrpc.HouseworkTemplatesResponse], error) {
	houseworkTemplates, err := houseworkUsecase.ListHouseworkTemplates(ctx, req.Msg.FamilyId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkTemplatesResponse{Templates: mapHouseworkTemplateListToGrpc(houseworkTemplates)})
	return res, nil
}

func (c *HouseworkController) CreateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.CreateHouseworkTemplate(ctx, mapHouseworkTemplateFromGrpc(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.UpdateHouseworkTemplate(ctx, mapHouseworkTemplateFromGrpc(req.Msg)); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkUsecase.DeleteHouseworkTemplate(ctx, req.Msg.Id); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.HouseworkPointRequest]) (*connect.Response[shwgrpc.HouseworkPointResponse], error) {
	houseworkPoint, err := houseworkUsecase.GetHouseworkPoint(ctx, req.Msg.UserId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkPointResponse{Point: mapHouseworkPointToGrpc(houseworkPoint)})
	return res, nil
}

func (c *HouseworkController) GetHouseworkPointHistory(ctx context.Context, req *connect.Request[shwgrpc.HouseworkPointHistoryRequest]) (*connect.Response[shwgrpc.HouseworkPointHistoryResponse], error) {
	houseworkPointHistories, err := houseworkUsecase.ListHouseworkPointHistory(ctx, req.Msg.UserId)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkPointHistoryResponse{History: mapHouseworkPointHistoryListToGrpc(houseworkPointHistories)})
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
		Status:   h.Status,
		WorkUser: &shwgrpc.UserInfo{
			Id:   h.WorkUser.ID,
			Name: h.WorkUser.Name,
		},
		StartedAt: h.StartedAt.Unix(),
		EndedAt:   h.EndedAt.Unix(),
	}
}

func mapHouseworkFromGrpc(h *shwgrpc.Housework) domain.Housework {
	return domain.Housework{
		ID:       h.Id,
		FamilyID: h.FamilyId,
		Title:    h.Title,
		Detail:   h.Detail,
		Status:   h.Status,
		WorkUser: domain.UserInfo{
			ID:   h.WorkUser.GetId(),
			Name: h.WorkUser.GetName(),
		},
		StartedAt: time.Unix(h.StartedAt, 0),
		EndedAt:   time.Unix(h.EndedAt, 0),
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

func mapHouseworkMemoFromGrpc(m *shwgrpc.HouseworkMemo) domain.HouseworkMemo {
	return domain.HouseworkMemo{
		ID:          m.Id,
		HouseworkID: m.HouseworkId,
		Message:     m.Message,
		SendFrom: domain.UserInfo{
			ID:   m.SendFrom.GetId(),
			Name: m.SendFrom.GetName(),
		},
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

func mapHouseworkTemplateFromGrpc(t *shwgrpc.HouseworkTemplate) domain.HouseworkTemplate {
	return domain.HouseworkTemplate{
		ID:       t.Id,
		FamilyID: t.FamilyId,
		Title:    t.Title,
		Detail:   t.Detail,
	}
}

func mapHouseworkTemplateListToGrpc(list []domain.HouseworkTemplate) []*shwgrpc.HouseworkTemplate {
	res := make([]*shwgrpc.HouseworkTemplate, 0, len(list))
	for _, item := range list {
		res = append(res, mapHouseworkTemplateToGrpc(item))
	}
	return res
}

func mapHouseworkPointToGrpc(point domain.HouseworkPoint) *shwgrpc.HouseworkPoint {
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
