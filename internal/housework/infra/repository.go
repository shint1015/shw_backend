package infra

import (
	"context"
	"shwgrpc/internal/housework/domain"
	"shwgrpc/internal/housework/port"
	"shwgrpc/model"
	"time"
)

type HouseworkRepository struct{}

func NewHouseworkRepository() port.HouseworkRepository {
	return &HouseworkRepository{}
}

func (r *HouseworkRepository) ListByFamilyID(_ context.Context, familyID uint64) ([]domain.Housework, error) {
	var res []model.Housework
	if err := model.DB.
		Where(&model.Housework{FamilyID: uint(familyID)}).
		Preload("WorkToUser").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return mapHouseworkList(res), nil
}

func (r *HouseworkRepository) GetDetail(_ context.Context, id uint64) (domain.Housework, []domain.HouseworkMemo, error) {
	var res model.Housework
	if err := model.DB.
		Where(&model.Housework{ID: uint(id)}).
		Preload("WorkToUser").
		Preload("Memo").
		Preload("Memo.SendFromUser").
		First(&res).Error; err != nil {
		return domain.Housework{}, nil, err
	}
	housework := mapHousework(res)
	memos := mapHouseworkMemoList(res.Memo)
	return housework, memos, nil
}

func (r *HouseworkRepository) Create(_ context.Context, housework domain.Housework) error {
	m := model.Housework{
		FamilyID:  uint(housework.FamilyID),
		Title:     housework.Title,
		Detail:    housework.Detail,
		Status:    housework.Status,
		WorkTo:    uint(housework.WorkUser.ID),
		StartedAt: housework.StartedAt,
		EndedAt:   housework.EndedAt,
	}
	return m.Create(nil)
}

func (r *HouseworkRepository) Update(_ context.Context, housework domain.Housework) error {
	m := model.Housework{
		ID:        uint(housework.ID),
		Title:     housework.Title,
		Detail:    housework.Detail,
		WorkTo:    uint(housework.WorkUser.ID),
		StartedAt: housework.StartedAt,
		EndedAt:   housework.EndedAt,
	}
	return m.Update(nil)
}

func (r *HouseworkRepository) UpdateStatus(_ context.Context, id uint64, status string) error {
	m := model.Housework{
		ID:     uint(id),
		Status: status,
	}
	return m.Update(nil)
}

func (r *HouseworkRepository) Delete(_ context.Context, id uint64) error {
	m := model.Housework{ID: uint(id)}
	return m.Delete(nil)
}

type HouseworkMemoRepository struct{}

func NewHouseworkMemoRepository() port.HouseworkMemoRepository {
	return &HouseworkMemoRepository{}
}

func (r *HouseworkMemoRepository) ListByHouseworkID(_ context.Context, houseworkID uint64) ([]domain.HouseworkMemo, error) {
	var res []model.HouseworkMemo
	if err := model.DB.
		Where(&model.HouseworkMemo{HouseworkID: uint(houseworkID)}).
		Preload("SendFromUser").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return mapHouseworkMemoList(&res), nil
}

func (r *HouseworkMemoRepository) Create(_ context.Context, memo domain.HouseworkMemo) error {
	m := model.HouseworkMemo{
		HouseworkID: uint(memo.HouseworkID),
		Message:     memo.Message,
		SendFrom:    uint(memo.SendFrom.ID),
	}
	return m.Create(nil)
}

func (r *HouseworkMemoRepository) Update(_ context.Context, memo domain.HouseworkMemo) error {
	m := model.HouseworkMemo{
		ID:      uint(memo.ID),
		Message: memo.Message,
	}
	return m.Update(nil)
}

func (r *HouseworkMemoRepository) Delete(_ context.Context, id uint64) error {
	m := model.HouseworkMemo{ID: uint(id)}
	return m.Delete(nil)
}

type HouseworkTemplateRepository struct{}

func NewHouseworkTemplateRepository() port.HouseworkTemplateRepository {
	return &HouseworkTemplateRepository{}
}

func (r *HouseworkTemplateRepository) Get(_ context.Context, id uint64) (domain.HouseworkTemplate, error) {
	m := model.HouseworkTemplate{ID: uint(id)}
	res, err := m.Get()
	if err != nil {
		return domain.HouseworkTemplate{}, err
	}
	return mapHouseworkTemplate(res), nil
}

func (r *HouseworkTemplateRepository) ListByFamilyID(_ context.Context, familyID uint64) ([]domain.HouseworkTemplate, error) {
	m := model.HouseworkTemplate{FamilyID: uint(familyID)}
	res, err := m.GetAll()
	if err != nil {
		return nil, err
	}
	return mapHouseworkTemplateList(res), nil
}

func (r *HouseworkTemplateRepository) Create(_ context.Context, template domain.HouseworkTemplate) error {
	m := model.HouseworkTemplate{
		FamilyID: uint(template.FamilyID),
		Title:    template.Title,
		Detail:   template.Detail,
	}
	return m.Create(nil)
}

func (r *HouseworkTemplateRepository) Update(_ context.Context, template domain.HouseworkTemplate) error {
	m := model.HouseworkTemplate{
		ID:     uint(template.ID),
		Title:  template.Title,
		Detail: template.Detail,
	}
	return m.Update(nil)
}

func (r *HouseworkTemplateRepository) Delete(_ context.Context, id uint64) error {
	m := model.HouseworkTemplate{ID: uint(id)}
	return m.Delete(nil)
}

type UserRepository struct{}

func NewUserRepository() port.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetWithHouseworkPoint(_ context.Context, id uint64) (domain.HouseworkPoint, error) {
	var res model.User
	if err := model.DB.Preload("HouseworkPoint").First(&res, id).Error; err != nil {
		return domain.HouseworkPoint{}, err
	}
	return mapHouseworkPoint(res), nil
}

type HouseworkPointHistoryRepository struct{}

func NewHouseworkPointHistoryRepository() port.HouseworkPointHistoryRepository {
	return &HouseworkPointHistoryRepository{}
}

func (r *HouseworkPointHistoryRepository) ListByUserSince(_ context.Context, userID uint64, since time.Time) ([]domain.HouseworkPointHistory, error) {
	var res []model.HouseworkPointHistory
	if err := model.DB.
		Where(&model.HouseworkPointHistory{UserID: uint(userID)}).
		Where("created_at > ?", since).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return mapHouseworkPointHistoryList(res), nil
}

func mapHouseworkList(list []model.Housework) []domain.Housework {
	res := make([]domain.Housework, 0, len(list))
	for _, item := range list {
		res = append(res, mapHousework(item))
	}
	return res
}

func mapHousework(h model.Housework) domain.Housework {
	return domain.Housework{
		ID:        uint64(h.ID),
		FamilyID:  uint64(h.FamilyID),
		Title:     h.Title,
		Detail:    h.Detail,
		Status:    h.Status,
		WorkUser:  mapUserInfo(h.WorkToUser),
		StartedAt: h.StartedAt,
		EndedAt:   h.EndedAt,
	}
}

func mapHouseworkMemoList(list *[]model.HouseworkMemo) []domain.HouseworkMemo {
	if list == nil {
		return nil
	}
	res := make([]domain.HouseworkMemo, 0, len(*list))
	for _, item := range *list {
		res = append(res, mapHouseworkMemo(item))
	}
	return res
}

func mapHouseworkMemo(m model.HouseworkMemo) domain.HouseworkMemo {
	return domain.HouseworkMemo{
		ID:          uint64(m.ID),
		HouseworkID: uint64(m.HouseworkID),
		Message:     m.Message,
		SendFrom:    mapUserInfo(m.SendFromUser),
	}
}

func mapHouseworkTemplate(t model.HouseworkTemplate) domain.HouseworkTemplate {
	return domain.HouseworkTemplate{
		ID:        uint64(t.ID),
		FamilyID:  uint64(t.FamilyID),
		Title:     t.Title,
		Detail:    t.Detail,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func mapHouseworkTemplateList(list []model.HouseworkTemplate) []domain.HouseworkTemplate {
	res := make([]domain.HouseworkTemplate, 0, len(list))
	for _, item := range list {
		res = append(res, mapHouseworkTemplate(item))
	}
	return res
}

func mapHouseworkPoint(u model.User) domain.HouseworkPoint {
	point := domain.HouseworkPoint{
		User: mapUserInfo(u),
	}
	if u.HouseworkPoint == nil {
		return point
	}
	point.Point = int64(u.HouseworkPoint.Point)
	point.CreatedAt = u.HouseworkPoint.CreatedAt
	point.UpdatedAt = u.HouseworkPoint.UpdatedAt
	return point
}

func mapHouseworkPointHistoryList(list []model.HouseworkPointHistory) []domain.HouseworkPointHistory {
	res := make([]domain.HouseworkPointHistory, 0, len(list))
	for _, item := range list {
		res = append(res, domain.HouseworkPointHistory{
			ID:        uint64(item.ID),
			Detail:    item.Detail,
			Point:     int64(item.Point),
			CreatedAt: item.CreatedAt,
		})
	}
	return res
}

func mapUserInfo(u model.User) domain.UserInfo {
	return domain.UserInfo{
		ID:   uint64(u.ID),
		Name: u.Name,
	}
}
