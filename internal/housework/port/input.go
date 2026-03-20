package port

import "time"

type CreateHouseworkInput struct {
	FamilyID   uint64
	Title      string
	Detail     string
	StatusID   uint64
	WorkUserID uint64
	StartedAt  time.Time
	EndedAt    time.Time
}

type UpdateHouseworkInput struct {
	ID         uint64
	Title      string
	Detail     string
	WorkUserID uint64
	StartedAt  time.Time
	EndedAt    time.Time
}

type CreateHouseworkMemoInput struct {
	HouseworkID uint64
	Message     string
	SendFromID  uint64
}

type UpdateHouseworkMemoInput struct {
	ID      uint64
	Message string
}

type CreateHouseworkTemplateInput struct {
	FamilyID uint64
	Title    string
	Detail   string
}

type UpdateHouseworkTemplateInput struct {
	ID     uint64
	Title  string
	Detail string
}
