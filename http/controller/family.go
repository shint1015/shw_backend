package controller

import (
	"context"
	shwgrpc "shwgrpc/pkg/grpc"
)

type FamilyController struct{}

func NewFamilyController() *FamilyController {
	return &FamilyController{}
}

func (c *FamilyController) GetFamily(ctx context.Context, req *shwgrpc.FamilyRequest) (*shwgrpc.FamilyResponse, error) {

}

func (c *FamilyController) CreateFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.CommonResponse, error) {

}

func (c *FamilyController) UpdateFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.CommonResponse, error) {

}

func (c *FamilyController) DeleteFamily(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.CommonResponse, error) {

}

func (c *FamilyController) AddFamilyMember(ctx context.Context, req *shwgrpc.Family) (*shwgrpc.CommonResponse, error) {

}
