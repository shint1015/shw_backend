package controller

import (
	"connectrpc.com/connect"
	"context"
	"shwgrpc/internal/service"
	shwgrpc "shwgrpc/pkg/grpc"
)

type HouseworkController struct{}

var houseworkService = service.NewHouseworkService()

func NewHouseworkController() *HouseworkController {
	return &HouseworkController{}
}

func (c *HouseworkController) GetHousework(ctx context.Context, req *connect.Request[shwgrpc.HouseworkRequest]) (*connect.Response[shwgrpc.HouseworkResponse], error) {
	houseworks, err := houseworkService.GetHousework(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkResponse{Housework: houseworks})
	return res, nil
}

func (c *HouseworkController) GetHouseworkDetail(ctx context.Context, req *connect.Request[shwgrpc.HouseworkDetailRequest]) (*connect.Response[shwgrpc.HouseworkDetailResponse], error) {
	housework, houseworkMemos, err := houseworkService.GetHouseworkDetail(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkDetailResponse{Housework: housework, Memo: houseworkMemos})
	return res, nil
}

func (c *HouseworkController) CreateHousework(ctx context.Context, req *connect.Request[shwgrpc.Housework]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.CreateHousework(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHousework(ctx context.Context, req *connect.Request[shwgrpc.Housework]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.UpdateHousework(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) FinishHousework(ctx context.Context, req *connect.Request[shwgrpc.Housework]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.FinishHousework(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHousework(ctx context.Context, req *connect.Request[shwgrpc.Housework]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.DeleteHousework(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemoRequest]) (*connect.Response[shwgrpc.HouseworkMemoResponse], error) {
	memoList, err := houseworkService.GetHouseworkMemo(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkMemoResponse{Memo: memoList})
	return res, nil
}

func (c *HouseworkController) CreateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.CreateHouseworkMemo(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.UpdateHouseworkMemo(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) DeleteHouseworkMemo(ctx context.Context, req *connect.Request[shwgrpc.HouseworkMemo]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.DeleteHouseworkMemo(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplateRequest]) (*connect.Response[shwgrpc.HouseworkTemplateResponse], error) {
	houseworkTemplates, err := houseworkService.GetHouseworkTemplate(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkTemplateResponse{Template: houseworkTemplates})
	return res, nil
}

func (c *HouseworkController) CreateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.CreateHouseworkTemplate(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) UpdateHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.UpdateHouseworkTemplate(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) RemoveHouseworkTemplate(ctx context.Context, req *connect.Request[shwgrpc.HouseworkTemplate]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := houseworkService.RemoveHouseworkTemplate(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}

func (c *HouseworkController) GetHouseworkPoint(ctx context.Context, req *connect.Request[shwgrpc.HouseworkPointRequest]) (*connect.Response[shwgrpc.HouseworkPointResponse], error) {
	houseworkPoints, err := houseworkService.GetHouseworkPoint(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkPointResponse{Point: houseworkPoints})
	return res, nil
}

func (c *HouseworkController) GetHouseworkPointHistory(ctx context.Context, req *connect.Request[shwgrpc.HouseworkPointHistoryRequest]) (*connect.Response[shwgrpc.HouseworkPointHistoryResponse], error) {
	houseworkPointHistories, err := houseworkService.GetHouseworkPointHistory(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.HouseworkPointHistoryResponse{History: houseworkPointHistories})
	return res, nil
}
