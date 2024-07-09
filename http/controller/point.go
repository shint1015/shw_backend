package controller

import (
	"connectrpc.com/connect"
	"context"
	"shwgrpc/internal/service"
	shwgrpc "shwgrpc/pkg/grpc"
)

type PointController struct{}

var pointService = service.NewPointService()

func NewPointController() *PointController {
	return &PointController{}
}

func (c *PointController) GetFamilyPointList(ctx context.Context, req *connect.Request[shwgrpc.FamilyPointListRequest]) (*connect.Response[shwgrpc.FamilyPointList], error) {
	points, err := pointService.GetFamilyPointList(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.FamilyPointList{Points: points})
	return res, nil
}
