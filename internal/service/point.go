package service

import (
	"shwgrpc/model"
	shwgrpc "shwgrpc/pkg/grpc"
)

type PointService struct{}

func NewPointService() *PointService {
	return &PointService{}
}

func (s *PointService) GetFamilyPointList(req *shwgrpc.FamilyPointListRequest) ([]*shwgrpc.Point, error) {
	familyId := uint(req.FamilyId)
	point := model.HouseworkPoint{
		FamilyID: familyId,
	}
	points, err := point.GetAll()
	if err != nil {
		return nil, err
	}
	var res []*shwgrpc.Point
	for _, val := range points {
		res = append(res, &shwgrpc.Point{
			Id:    uint64(val.ID),
			Point: int64(val.Point),
			User: &shwgrpc.UserInfo{
				Id:   uint64(val.UserID),
				Name: val.User.Name,
			},
		})
	}
	return res, nil
}
