package service

import (
	"shwgrpc/model"
	shwgrpc "shwgrpc/pkg/grpc"
	"time"
)

type HouseworkService struct{}
type HouseworkStatus string

const (
	HouseworkStatusPlan  = HouseworkStatus("plan")
	HouseworkStatusDoing = HouseworkStatus("doing")
	HouseworkStatusDone  = HouseworkStatus("done")
)

func NewHouseworkService() *HouseworkService {
	return &HouseworkService{}
}

func (s *HouseworkService) GetHousework(req *shwgrpc.HouseworkRequest) ([]*shwgrpc.Housework, error) {
	housework := model.Housework{
		FamilyID: uint(req.FamilyId),
	}
	houseworks, err := housework.GetAll()

	if err != nil {
		return nil, err
	}
	var res []*shwgrpc.Housework
	for _, val := range houseworks {
		res = append(res, s.createFormatGrpcHousework(val))
	}
	return res, nil
}

func (s *HouseworkService) GetHouseworkDetail(req *shwgrpc.HouseworkDetailRequest) (*shwgrpc.Housework, []*shwgrpc.HouseworkMemo, error) {
	houseworkModel := model.Housework{
		ID: uint(req.Id),
	}
	housework, err := houseworkModel.Get()
	if err != nil {
		return nil, nil, err
	}
	resHousework := s.createFormatGrpcHousework(*housework)
	var resHouseworkMemo []*shwgrpc.HouseworkMemo
	for _, val := range housework.Memo {
		resHouseworkMemo = append(resHouseworkMemo, s.createFormatGrpcHouseworkMemo(val))
	}
	return resHousework, resHouseworkMemo, nil
}

func (s *HouseworkService) CreateHousework(req *shwgrpc.Housework) error {
	// TODO: familyIdから存在するかどうか
	// TODO: login情報から、所属するfamilyIdを取得し、それ以外のfamilyIdを指定されたらエラー
	housework := model.Housework{
		FamilyID:  uint(req.FamilyId),
		Title:     req.Title,
		Detail:    req.Detail,
		Status:    string(HouseworkStatusPlan),
		WorkTo:    uint(req.WorkUser.Id),
		StartedAt: time.Unix(req.StartedAt, 0),
		EndedAt:   time.Unix(req.EndedAt, 0),
	}
	if err := housework.Create(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) UpdateHousework(req *shwgrpc.Housework) error {
	// TODO: familyIdから存在するかどうか
	// TODO: login情報から、所属するfamilyIdを取得し、それ以外のfamilyIdを指定されたらエラー
	housework := model.Housework{
		ID:        uint(req.Id),
		Title:     req.Title,
		Detail:    req.Detail,
		WorkTo:    uint(req.WorkUser.Id),
		StartedAt: time.Unix(req.StartedAt, 0),
		EndedAt:   time.Unix(req.EndedAt, 0),
	}
	if err := housework.Update(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) FinishHousework(req *shwgrpc.Housework) error {
	// TODO: familyIdから存在するかどうか
	// TODO: login情報から、所属するfamilyIdを取得し、それ以外のfamilyIdを指定されたらエラー
	housework := model.Housework{
		ID:      uint(req.Id),
		WorkTo:  uint(req.WorkUser.Id),
		Status:  string(HouseworkStatusDone),
		EndedAt: time.Unix(req.EndedAt, 0),
	}
	if err := housework.Update(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) DeleteHousework(req *shwgrpc.Housework) error {
	// TODO: familyIdから存在するかどうか
	// TODO: login情報から、所属するfamilyIdを取得し、それ以外のfamilyIdを指定されたらエラー

	housework := model.Housework{
		ID: uint(req.Id),
	}
	if err := housework.Delete(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) GetHouseworkMemo(req *shwgrpc.HouseworkMemoRequest) ([]*shwgrpc.HouseworkMemo, error) {
	// TODO: familyIdから存在するかどうか
	// TODO: login情報から、所属するfamilyIdを取得し、それ以外のfamilyIdを指定されたらエラー

	houseworkMemo := model.HouseworkMemo{
		HouseworkID: uint(req.HouseworkId),
	}
	houseworkMemos, err := houseworkMemo.GetAll()
	if err != nil {
		return nil, err
	}
	var res []*shwgrpc.HouseworkMemo
	for _, val := range houseworkMemos {
		res = append(res, s.createFormatGrpcHouseworkMemo(val))
	}
	return res, nil
}

func (s *HouseworkService) CreateHouseworkMemo(req *shwgrpc.HouseworkMemo) error {
	// TODO: familyIdから存在するかどうか
	// TODO: login情報から、所属するfamilyIdを取得し、それ以外のfamilyIdを指定されたらエラー
	// TODO: houseworkIdから存在するかどうか
	userId := uint(1)
	houseworkMemo := model.HouseworkMemo{
		HouseworkID: uint(req.HouseworkId),
		Message:     req.Message,
		SendFrom:    userId,
	}

	if err := houseworkMemo.Create(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) UpdateHouseworkMemo(req *shwgrpc.HouseworkMemo) error {
	// TODO: familyIdから存在するかどうか
	// TODO: login情報から、所属するfamilyIdを取得し、それ以外のfamilyIdを指定されたらエラー
	// TODO: houseworkIdから存在するかどうか
	houseworkMemo := model.HouseworkMemo{
		ID:      uint(req.Id),
		Message: req.Message,
	}
	if err := houseworkMemo.Update(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) DeleteHouseworkMemo(req *shwgrpc.HouseworkMemo) error {
	// TODO: familyIdから存在するかどうか
	// TODO: login情報から、所属するfamilyIdを取得し、それ以外のfamilyIdを指定されたらエラー
	// TODO: houseworkIdから存在するかどうか
	houseworkMemo := model.HouseworkMemo{
		ID: uint(req.Id),
	}
	if err := houseworkMemo.Delete(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) GetHouseworkPoint(req *shwgrpc.HouseworkPointRequest) (*shwgrpc.HouseworkPoint, error) {
	u := model.User{
		ID: uint(req.UserId),
	}
	user, err := u.Get()
	if err != nil {
		return nil, err
	}
	return s.createFormatGrpcHouseworkPoint(*user), nil
}

func (s *HouseworkService) GetHouseworkPointHistory(req *shwgrpc.HouseworkPointHistoryRequest) ([]*shwgrpc.HouseworkPointHistory, error) {
	hwh := model.HouseworkPointHistory{
		UserID: uint(req.UserId),
	}
	target := time.Now()
	target.AddDate(0, 0, -7)
	where := [][]interface{}{{"created_at > ", target}}
	houseworkHistories, err := hwh.GetAll(where)
	if err != nil {
		return nil, err
	}
	var res []*shwgrpc.HouseworkPointHistory
	for _, val := range houseworkHistories {
		res = append(res, s.createFormatGrpcHouseworkPointHistory(val))
	}
	return res, nil
}

func (s *HouseworkService) GetHouseworkTemplate(req *shwgrpc.HouseworkTemplateRequest) ([]*shwgrpc.HouseworkTemplate, error) {
	houseworkTemplate := model.HouseworkTemplate{
		FamilyID: uint(req.FamilyId),
	}
	houseworkTemplates, err := houseworkTemplate.GetAll()
	if err != nil {
		return nil, err
	}
	var res []*shwgrpc.HouseworkTemplate
	for _, val := range houseworkTemplates {
		res = append(res, s.createFormatGrpcHouseworkTemplate(val))
	}
	return res, nil
}

func (s *HouseworkService) CreateHouseworkTemplate(req *shwgrpc.HouseworkTemplate) error {
	houseworkTemplate := model.HouseworkTemplate{
		FamilyID: uint(req.FamilyId),
		Title:    req.Title,
		Detail:   req.Detail,
	}
	if err := houseworkTemplate.Create(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) UpdateHouseworkTemplate(req *shwgrpc.HouseworkTemplate) error {
	houseworkTemplate := model.HouseworkTemplate{
		ID:     uint(req.Id),
		Title:  req.Title,
		Detail: req.Detail,
	}
	if err := houseworkTemplate.Update(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) RemoveHouseworkTemplate(req *shwgrpc.HouseworkTemplate) error {
	houseworkTemplate := model.HouseworkTemplate{
		ID: uint(req.Id),
	}
	if err := houseworkTemplate.Delete(nil); err != nil {
		return err
	}
	return nil
}

func (s *HouseworkService) createFormatGrpcHousework(housework model.Housework) *shwgrpc.Housework {
	workToUser := housework.WorkToUser
	return &shwgrpc.Housework{
		Id:       uint64(housework.ID),
		FamilyId: uint64(housework.FamilyID),
		Title:    housework.Title,
		Detail:   housework.Detail,
		Status:   housework.Status,
		WorkUser: &shwgrpc.UserInfo{
			Id:   uint64(workToUser.ID),
			Name: workToUser.Name,
		},
		StartedAt: housework.StartedAt.Unix(),
		EndedAt:   housework.EndedAt.Unix(),
	}
}

func (s *HouseworkService) createFormatGrpcHouseworkMemo(memo model.HouseworkMemo) *shwgrpc.HouseworkMemo {
	sendFromUser := memo.SendFromUser
	return &shwgrpc.HouseworkMemo{
		Id:          uint64(memo.ID),
		HouseworkId: uint64(memo.HouseworkID),
		Message:     memo.Message,
		DraftUser: &shwgrpc.UserInfo{
			Id:   uint64(sendFromUser.ID),
			Name: sendFromUser.Name,
		},
	}
}

func (s *HouseworkService) createFormatGrpcHouseworkPoint(user model.User) *shwgrpc.HouseworkPoint {
	return &shwgrpc.HouseworkPoint{
		Point: int64(user.HouseworkPoint.Point),
		User: &shwgrpc.UserInfo{
			Id:   uint64(user.ID),
			Name: user.Name,
		},
		CreatedAt: user.HouseworkPoint.CreatedAt.Unix(),
		UpdatedAt: user.HouseworkPoint.UpdatedAt.Unix(),
	}
}

func (s *HouseworkService) createFormatGrpcHouseworkPointHistory(history model.HouseworkPointHistory) *shwgrpc.HouseworkPointHistory {
	return &shwgrpc.HouseworkPointHistory{
		Id:        uint64(history.ID),
		Detail:    history.Detail,
		Point:     int64(history.Point),
		CreatedAt: history.CreatedAt.Unix(),
	}
}

func (s *HouseworkService) createFormatGrpcHouseworkTemplate(template model.HouseworkTemplate) *shwgrpc.HouseworkTemplate {
	return &shwgrpc.HouseworkTemplate{
		Id:        uint64(template.ID),
		FamilyId:  uint64(template.FamilyID),
		Title:     template.Title,
		Detail:    template.Detail,
		CreatedAt: template.CreatedAt.Unix(),
		UpdatedAt: template.UpdatedAt.Unix(),
	}
}
