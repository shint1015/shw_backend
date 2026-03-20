package domain

import "time"

type HouseworkStatus string

const (
	HouseworkStatusPlan  HouseworkStatus = "plan"
	HouseworkStatusDoing HouseworkStatus = "doing"
	HouseworkStatusDone  HouseworkStatus = "done"
)

const (
	HouseworkStatusIDUnspecified uint64 = 0
	HouseworkStatusIDPlan        uint64 = 1
	HouseworkStatusIDDoing       uint64 = 2
	HouseworkStatusIDDone        uint64 = 3
)

func HouseworkStatusFromID(id uint64) HouseworkStatus {
	switch id {
	case HouseworkStatusIDPlan:
		return HouseworkStatusPlan
	case HouseworkStatusIDDoing:
		return HouseworkStatusDoing
	case HouseworkStatusIDDone:
		return HouseworkStatusDone
	default:
		return ""
	}
}

func HouseworkStatusToID(status HouseworkStatus) uint64 {
	switch status {
	case HouseworkStatusPlan:
		return HouseworkStatusIDPlan
	case HouseworkStatusDoing:
		return HouseworkStatusIDDoing
	case HouseworkStatusDone:
		return HouseworkStatusIDDone
	default:
		return HouseworkStatusIDUnspecified
	}
}

type UserInfo struct {
	ID   uint64
	Name string
}

type Housework struct {
	ID        uint64
	FamilyID  uint64
	Title     string
	Detail    string
	StatusID  uint64
	Status    HouseworkStatus
	WorkUser  UserInfo
	StartedAt time.Time
	EndedAt   time.Time
}

type HouseworkMemo struct {
	ID          uint64
	HouseworkID uint64
	Message     string
	SendFrom    UserInfo
}

type HouseworkPoint struct {
	User      UserInfo
	Point     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HouseworkPointHistory struct {
	ID        uint64
	Detail    string
	Point     int64
	CreatedAt time.Time
}

type HouseworkTemplate struct {
	ID        uint64
	FamilyID  uint64
	Title     string
	Detail    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
