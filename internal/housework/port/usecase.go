package port

import (
	"context"
	"shwgrpc/internal/housework/domain"
)

type HouseworkUsecase interface {
	ListHousework(ctx context.Context, familyID uint64) ([]domain.Housework, error)
	GetHouseworkDetail(ctx context.Context, id uint64) (domain.Housework, []domain.HouseworkMemo, error)
	CreateHousework(ctx context.Context, housework domain.Housework) error
	UpdateHousework(ctx context.Context, housework domain.Housework) error
	FinishHousework(ctx context.Context, id uint64) error
	DeleteHousework(ctx context.Context, id uint64) error

	ListHouseworkMemo(ctx context.Context, houseworkID uint64) ([]domain.HouseworkMemo, error)
	CreateHouseworkMemo(ctx context.Context, memo domain.HouseworkMemo) error
	UpdateHouseworkMemo(ctx context.Context, memo domain.HouseworkMemo) error
	DeleteHouseworkMemo(ctx context.Context, id uint64) error

	GetHouseworkPoint(ctx context.Context, userID uint64) (domain.HouseworkPoint, error)
	ListHouseworkPointHistory(ctx context.Context, userID uint64) ([]domain.HouseworkPointHistory, error)

	GetHouseworkTemplate(ctx context.Context, id uint64) (domain.HouseworkTemplate, error)
	ListHouseworkTemplates(ctx context.Context, familyID uint64) ([]domain.HouseworkTemplate, error)
	CreateHouseworkTemplate(ctx context.Context, template domain.HouseworkTemplate) error
	UpdateHouseworkTemplate(ctx context.Context, template domain.HouseworkTemplate) error
	DeleteHouseworkTemplate(ctx context.Context, id uint64) error
}
