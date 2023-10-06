package controller

import (
	"context"
	shwgrpc "shwgrpc/pkg/grpc"
)

type HouseworkController struct{}

func NewHouseworkController() *HouseworkController {
	return &HouseworkController{}
}

func (c *HouseworkController) GetHousework(ctx context.Context, req *shwgrpc.HouseworkRequest) (*shwgrpc.HouseworkResponse, error) {

}

func (c *HouseworkController) GetHouseworkDetail(ctx context.Context, req *shwgrpc.HouseworkRequest) (*shwgrpc.HouseworkResponse, error) {
}

func (c *HouseworkController) CreateHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.CommonResponse, error) {
}

func (c *HouseworkController) UpdateHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.CommonResponse, error) {
}

func (c *HouseworkController) FinishHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.CommonResponse, error) {
}

func (c *HouseworkController) DeleteHousework(ctx context.Context, req *shwgrpc.Housework) (*shwgrpc.CommonResponse, error) {
}

func (c *HouseworkController) GetHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkDetailRequest) (*shwgrpc.HouseworkDetailResponse, error) {
}

func (c *HouseworkController) CreateHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.CommonResponse, error) {
}

func (c *HouseworkController) UpdateHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.CommonResponse, error) {
}

func (c *HouseworkController) DeleteHouseworkMemo(ctx context.Context, req *shwgrpc.HouseworkMemo) (*shwgrpc.CommonResponse, error) {
}

func (c *HouseworkController) GetHouseworkPoint(ctx context.Context, req *shwgrpc.HouseworkPointRequest) (*shwgrpc.HouseworkPointResponse, error) {
}

func (c *HouseworkController) GetHouseworkPointHistory(ctx context.Context, req *shwgrpc.HouseworkPointHistoryRequest) (*shwgrpc.HouseworkPointHistoryResponse, error) {
}
