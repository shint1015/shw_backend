package controller

import (
	"context"
	"shwgrpc/internal/service"
	shwgrpc "shwgrpc/pkg/grpc"
)

type UserController struct{}

var userService = service.NewUserService()

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) UpdateRole(ctx context.Context, req *shwgrpc.UpdateRoleRequest) (*shwgrpc.CommonResponse, error) {
	if err := userService.UpdateRole(req); err != nil {
		return nil, err
	}
	return &shwgrpc.CommonResponse{Message: "success"}, nil
}
