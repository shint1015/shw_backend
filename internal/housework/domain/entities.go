package domain

import "time"

type HouseworkStatus string

const (
	HouseworkStatusPlan  HouseworkStatus = "plan"
	HouseworkStatusDoing HouseworkStatus = "doing"
	HouseworkStatusDone  HouseworkStatus = "done"
)

type UserInfo struct {
	ID   uint64
	Name string
}

type Housework struct {
	ID        uint64
	FamilyID  uint64
	Title     string
	Detail    string
	Status    string
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
