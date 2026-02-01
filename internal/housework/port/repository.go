package port

import (
	"context"
	"shwgrpc/internal/housework/domain"
	"time"
)

type HouseworkRepository interface {
	ListByFamilyID(ctx context.Context, familyID uint64) ([]domain.Housework, error)
	GetDetail(ctx context.Context, id uint64) (domain.Housework, []domain.HouseworkMemo, error)
	Create(ctx context.Context, housework domain.Housework) error
	Update(ctx context.Context, housework domain.Housework) error
	UpdateStatus(ctx context.Context, id uint64, status string) error
	Delete(ctx context.Context, id uint64) error
}

type HouseworkMemoRepository interface {
	ListByHouseworkID(ctx context.Context, houseworkID uint64) ([]domain.HouseworkMemo, error)
	Create(ctx context.Context, memo domain.HouseworkMemo) error
	Update(ctx context.Context, memo domain.HouseworkMemo) error
	Delete(ctx context.Context, id uint64) error
}

type HouseworkTemplateRepository interface {
	Get(ctx context.Context, id uint64) (domain.HouseworkTemplate, error)
	ListByFamilyID(ctx context.Context, familyID uint64) ([]domain.HouseworkTemplate, error)
	Create(ctx context.Context, template domain.HouseworkTemplate) error
	Update(ctx context.Context, template domain.HouseworkTemplate) error
	Delete(ctx context.Context, id uint64) error
}

type UserRepository interface {
	GetWithHouseworkPoint(ctx context.Context, id uint64) (domain.HouseworkPoint, error)
}

type HouseworkPointHistoryRepository interface {
	ListByUserSince(ctx context.Context, userID uint64, since time.Time) ([]domain.HouseworkPointHistory, error)
}
