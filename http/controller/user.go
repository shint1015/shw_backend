package controller

import (
	"connectrpc.com/connect"
	"context"
	"shwgrpc/internal/service"
	shwgrpc "shwgrpc/pkg/grpc"
)

type UserController struct{}

var userService = service.NewUserService()

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) UpdateRole(ctx context.Context, req *connect.Request[shwgrpc.UpdateRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := userService.UpdateRole(req.Msg); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}
