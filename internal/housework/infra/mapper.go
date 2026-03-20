package infra

import (
	"shwgrpc/internal/housework/domain"
	sharedmapper "shwgrpc/internal/shared/mapper"
	"shwgrpc/model"
)

func mapHouseworkList(list []model.Housework) []domain.Housework {
	res := make([]domain.Housework, 0, len(list))
	for _, item := range list {
		res = append(res, mapHousework(item))
	}
	return res
}
func mapUserInfo(u model.User) domain.UserInfo {
	return sharedmapper.MapUserInfoAs(u, func(id uint64, name string) domain.UserInfo {
		return domain.UserInfo{
			ID:   id,
			Name: name,
		}
	})
}

func mapHousework(h model.Housework) domain.Housework {
	status := domain.HouseworkStatus(h.Status)
	return domain.Housework{
		ID:        uint64(h.ID),
		FamilyID:  uint64(h.FamilyID),
		Title:     h.Title,
		Detail:    h.Detail,
		StatusID:  domain.HouseworkStatusToID(status),
		Status:    status,
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
