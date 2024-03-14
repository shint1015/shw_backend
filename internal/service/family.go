package service

import (
	"shwgrpc/model"
	shwgrpc "shwgrpc/pkg/grpc"
	"time"
)

type FamilyService struct{}

func NewFamilyService() *FamilyService {
	return &FamilyService{}
}

func (s *FamilyService) GetFamily(req *shwgrpc.FamilyRequest) (*shwgrpc.Family, error) {
	f := model.Family{
		ID: uint(req.FamilyId),
	}
	family, err := f.Get()
	if err != nil {
		return nil, err
	}
	return s.createFormatGrpcFamily(*family), nil
}

func (s *FamilyService) CreateFamily(req *shwgrpc.Family) error {
	//TODO: requestしたuserを取得
	userId := 1

	tx := model.DB.Begin()
	f := model.Family{
		Name:             req.Name,
		PointPerWorkTime: req.PointPerWorkTime,
	}
	if err := f.Create(tx); err != nil {
		tx.Rollback()
		return err
	}
	u := model.User{
		ID:       uint(userId),
		FamilyID: &f.ID,
	}
	if err := u.Update(tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *FamilyService) UpdateFamily(req *shwgrpc.Family) error {
	f := model.Family{
		ID:               uint(req.Id),
		Name:             req.Name,
		PointPerWorkTime: req.PointPerWorkTime,
	}
	if err := f.Update(nil); err != nil {
		return err
	}
	return nil
}

func (s *FamilyService) DeleteFamily(req *shwgrpc.Family) error {
	f := model.Family{
		ID: uint(req.Id),
	}
	if err := f.Delete(nil); err != nil {
		return err
	}
	return nil
}

func (s *FamilyService) GetFamilyHouseworkPoints(req *shwgrpc.FamilyHouseworkPointRequest) ([]*shwgrpc.FamilyHouseworkPoint, error) {
	familyId := uint(req.FamilyId)
	u := model.User{
		FamilyID: &familyId,
	}
	users, err := u.GetUsersByFamilyID()
	if err != nil {
		return nil, err
	}
	var res []*shwgrpc.FamilyHouseworkPoint
	for _, val := range users {
		res = append(res, s.createFormatGrpcHouseworkPoint(val))
	}
	return res, nil
}

func (s *FamilyService) AddFamilyMember(req *shwgrpc.AddFamilyMemberRequest) error {
	//TODO: ユーザの承認フローを入れる（メール => メールから承認 => OK）
	u := model.User{
		Email: req.Email,
		Name:  req.Name,
	}
	user, err := u.Get()
	if err != nil {
		return err
	}
	now := time.Now()
	expiredTime := now.Add(6 * time.Hour)

	familyId := uint(req.FamilyId)
	user.FamilyID = &familyId
	user.IsFamilyVerified = false
	user.FamilyVerifyExpireAt = &expiredTime
	if err := user.Update(nil); err != nil {
		return err
	}
	return nil
}

func (s *FamilyService) VerifyAddFamily() error {
	//TODO: URLからTokenが妥当かを確認？？
	return nil
}

func (s *FamilyService) createFormatGrpcFamily(template model.Family) *shwgrpc.Family {
	var familyMember []*shwgrpc.UserInfo
	for _, val := range *template.Users {
		familyMember = append(familyMember, &shwgrpc.UserInfo{
			Id:   uint64(val.ID),
			Name: val.Name,
			Role: val.Role.Name,
		})
	}
	return &shwgrpc.Family{
		Id:               uint64(template.ID),
		User:             familyMember,
		Name:             template.Name,
		PointPerWorkTime: template.PointPerWorkTime,
	}
}

func (s *FamilyService) createFormatGrpcHouseworkPoint(user model.User) *shwgrpc.FamilyHouseworkPoint {
	return &shwgrpc.FamilyHouseworkPoint{
		Point: int64(user.HouseworkPoint.Point),
		User: &shwgrpc.UserInfo{
			Id:   uint64(user.ID),
			Name: user.Name,
		},
		CreatedAt: user.HouseworkPoint.CreatedAt.Unix(),
		UpdatedAt: user.HouseworkPoint.UpdatedAt.Unix(),
	}
}
