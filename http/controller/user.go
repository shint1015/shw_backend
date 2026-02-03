package controller

import (
	"connectrpc.com/connect"
	"context"
	"shwgrpc/internal/user/infra"
	"shwgrpc/internal/user/port"
	"shwgrpc/internal/user/usecase"
	shwgrpc "shwgrpc/pkg/grpc"
)

type UserController struct{}

var userUsecase port.UserUsecase = usecase.NewUserUsecase(
	infra.NewUserRepository(),
)

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) UpdateRole(ctx context.Context, req *connect.Request[shwgrpc.UpdateRoleRequest]) (*connect.Response[shwgrpc.CommonResponse], error) {
	if err := userUsecase.UpdateRole(ctx, req.Msg.UserId, req.Msg.RoleId); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&shwgrpc.CommonResponse{Message: "success"})
	return res, nil
}
