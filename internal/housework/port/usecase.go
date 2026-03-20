package port

import (
	"context"
	"shwgrpc/internal/housework/domain"
)

type HouseworkUsecase interface {
	ListHousework(ctx context.Context, familyID uint64) ([]domain.Housework, error)
	GetHouseworkDetail(ctx context.Context, id uint64) (domain.Housework, []domain.HouseworkMemo, error)
	CreateHousework(ctx context.Context, input CreateHouseworkInput) error
	UpdateHousework(ctx context.Context, input UpdateHouseworkInput) error
	FinishHousework(ctx context.Context, id uint64) error
	DeleteHousework(ctx context.Context, id uint64) error

	ListHouseworkMemo(ctx context.Context, houseworkID uint64) ([]domain.HouseworkMemo, error)
	CreateHouseworkMemo(ctx context.Context, input CreateHouseworkMemoInput) error
	UpdateHouseworkMemo(ctx context.Context, input UpdateHouseworkMemoInput) error
	DeleteHouseworkMemo(ctx context.Context, id uint64) error

	GetHouseworkPoint(ctx context.Context, userID uint64) (domain.HouseworkPoint, error)
	ListHouseworkPointHistory(ctx context.Context, userID uint64) ([]domain.HouseworkPointHistory, error)

	GetHouseworkTemplate(ctx context.Context, id uint64) (domain.HouseworkTemplate, error)
	ListHouseworkTemplates(ctx context.Context, familyID uint64) ([]domain.HouseworkTemplate, error)
	CreateHouseworkTemplate(ctx context.Context, input CreateHouseworkTemplateInput) error
	UpdateHouseworkTemplate(ctx context.Context, input UpdateHouseworkTemplateInput) error
	DeleteHouseworkTemplate(ctx context.Context, id uint64) error
}
