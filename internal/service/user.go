package service

import (
	"shwgrpc/model"
	shwgrpc "shwgrpc/pkg/grpc"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) UpdateRole(req *shwgrpc.UpdateRoleRequest) error {
	//TODO: userIDからuserを取得
	//TODO: login情報とuser情報の比較
	//TODO: roleが存在するかどうか
	roleId := uint(req.RoleId)
	u := model.User{
		ID:     uint(req.UserId),
		RoleID: &roleId,
	}
	if err := u.Update(nil); err != nil {
		return err
	}
	
	return nil
}
