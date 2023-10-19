package controller

import (
	"context"
	"shwgrpc/internal/service"
	shwgrpc "shwgrpc/pkg/grpc"
)

type HouseworkController struct{}

var houseworkService = service.NewHouseworkService()

func NewHouseworkController() *HouseworkController {
	return &HouseworkController{}
}

func (c *HouseworkController) GetHousework(ctx context.Context, req *shwgrpc.HouseworkRequest) (*shwgrpc.HouseworkResponse, error) {
	houseworks, err := houseworkService.GetHousework(req)
	if err != nil {
		return nil, err
	}
	return &shwgrpc.HouseworkResponse{Housework: houseworks}, nil
}

func (c *HouseworkController) GetHouseworkDetail(ctx context.Context, req *shwgrpc.HouseworkDetailRequest) (*shwgrpc.HouseworkDetailResponse, error) {
	housework, houseworkMemos, err := houseworkService.GetHouseworkDetail(req)
	if err != nil {
		return nil, err
	}
	return &shwgrpc.HouseworkDetailResponse{Housework: housework, Memo: houseworkMemos}, nil
}

func (c *HouseworkController) CreateHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.CreateHousework(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) UpdateHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.UpdateHousework(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) FinishHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.FinishHousework(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) DeleteHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.DeleteHousework(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) GetHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemoRequest) (*shwgrpc.HouseworkMemoResponse, error) {
	memoList, err := houseworkService.GetHouseworkMemo(req)
	if err != nil {
		return nil, err
	}
	return &shwgrpc.HouseworkMemoResponse{Memo: memoList}, nil
}

func (c *HouseworkController) CreateHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.CreateHouseworkMemo(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) UpdateHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.UpdateHouseworkMemo(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) DeleteHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.DeleteHouseworkMemo(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) GetHouseworkTemplate(ctx context.Context, req *shwgrpc.HouseworkTemplateRequest) (*shwgrpc.HouseworkTemplateResponse, error) {
	houseworkTemplates, err := houseworkService.GetHouseworkTemplate(req)
	if err != nil {
		return nil, err
	}
	return &shwgrpc.HouseworkTemplateResponse{Template: houseworkTemplates}, nil
}

func (c *HouseworkController) CreateHouseworkTemplate(ctx context.Context, req *shwgrpc.HouseworkTemplate) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.CreateHouseworkTemplate(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) UpdateHouseworkTemplate(ctx context.Context, req *shwgrpc.HouseworkTemplate) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.UpdateHouseworkTemplate(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) RemoveHouseworkTemplate(ctx context.Context, req *shwgrpc.HouseworkTemplate) (*shwgrpc.CommonResponse, error) {
	if err := houseworkService.RemoveHouseworkTemplate(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}

func (c *HouseworkController) GetHouseworkPoint(ctx context.Context, req *shwgrpc.HouseworkPointRequest) (*shwgrpc.HouseworkPointResponse, error) {
	houseworkPoints, err := houseworkService.GetHouseworkPoint(req)
	if err != nil {
		return nil, err
	}
	return &shwgrpc.HouseworkPointResponse{Point: houseworkPoints}, nil
}

func (c *HouseworkController) GetHouseworkPointHistory(ctx context.Context, req *shwgrpc.HouseworkPointHistoryRequest) (*shwgrpc.HouseworkPointHistoryResponse, error) {
	houseworkPointHistories, err := houseworkService.GetHouseworkPointHistory(req)
	if err != nil {
		return nil, err
	}
	return &shwgrpc.HouseworkPointHistoryResponse{History: houseworkPointHistories}, nil
}
