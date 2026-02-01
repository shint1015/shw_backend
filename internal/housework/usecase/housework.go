package usecase

import (
	"context"
	"shwgrpc/internal/housework/domain"
	"shwgrpc/internal/housework/port"
	"time"
)

type HouseworkUsecase struct {
	houseworkRepo             port.HouseworkRepository
	houseworkMemoRepo         port.HouseworkMemoRepository
	houseworkTemplateRepo     port.HouseworkTemplateRepository
	userRepo                  port.UserRepository
	houseworkPointHistoryRepo port.HouseworkPointHistoryRepository
}

func NewHouseworkUsecase(
	houseworkRepo port.HouseworkRepository,
	houseworkMemoRepo port.HouseworkMemoRepository,
	houseworkTemplateRepo port.HouseworkTemplateRepository,
	userRepo port.UserRepository,
	houseworkPointHistoryRepo port.HouseworkPointHistoryRepository,
) *HouseworkUsecase {
	return &HouseworkUsecase{
		houseworkRepo:             houseworkRepo,
		houseworkMemoRepo:         houseworkMemoRepo,
		houseworkTemplateRepo:     houseworkTemplateRepo,
		userRepo:                  userRepo,
		houseworkPointHistoryRepo: houseworkPointHistoryRepo,
	}
}

func (u *HouseworkUsecase) ListHousework(ctx context.Context, familyID uint64) ([]domain.Housework, error) {
	return u.houseworkRepo.ListByFamilyID(ctx, familyID)
}

func (u *HouseworkUsecase) GetHouseworkDetail(ctx context.Context, id uint64) (domain.Housework, []domain.HouseworkMemo, error) {
	return u.houseworkRepo.GetDetail(ctx, id)
}

func (u *HouseworkUsecase) CreateHousework(ctx context.Context, housework domain.Housework) error {
	housework.Status = string(domain.HouseworkStatusPlan)
	return u.houseworkRepo.Create(ctx, housework)
}

func (u *HouseworkUsecase) UpdateHousework(ctx context.Context, housework domain.Housework) error {
	return u.houseworkRepo.Update(ctx, housework)
}

func (u *HouseworkUsecase) FinishHousework(ctx context.Context, id uint64) error {
	return u.houseworkRepo.UpdateStatus(ctx, id, string(domain.HouseworkStatusDone))
}

func (u *HouseworkUsecase) DeleteHousework(ctx context.Context, id uint64) error {
	return u.houseworkRepo.Delete(ctx, id)
}

func (u *HouseworkUsecase) ListHouseworkMemo(ctx context.Context, houseworkID uint64) ([]domain.HouseworkMemo, error) {
	return u.houseworkMemoRepo.ListByHouseworkID(ctx, houseworkID)
}

func (u *HouseworkUsecase) CreateHouseworkMemo(ctx context.Context, memo domain.HouseworkMemo) error {
	return u.houseworkMemoRepo.Create(ctx, memo)
}

func (u *HouseworkUsecase) UpdateHouseworkMemo(ctx context.Context, memo domain.HouseworkMemo) error {
	return u.houseworkMemoRepo.Update(ctx, memo)
}

func (u *HouseworkUsecase) DeleteHouseworkMemo(ctx context.Context, id uint64) error {
	return u.houseworkMemoRepo.Delete(ctx, id)
}

func (u *HouseworkUsecase) GetHouseworkPoint(ctx context.Context, userID uint64) (domain.HouseworkPoint, error) {
	return u.userRepo.GetWithHouseworkPoint(ctx, userID)
}

func (u *HouseworkUsecase) ListHouseworkPointHistory(ctx context.Context, userID uint64) ([]domain.HouseworkPointHistory, error) {
	since := time.Now().AddDate(0, 0, -7)
	return u.houseworkPointHistoryRepo.ListByUserSince(ctx, userID, since)
}

func (u *HouseworkUsecase) GetHouseworkTemplate(ctx context.Context, id uint64) (domain.HouseworkTemplate, error) {
	return u.houseworkTemplateRepo.Get(ctx, id)
}

func (u *HouseworkUsecase) ListHouseworkTemplates(ctx context.Context, familyID uint64) ([]domain.HouseworkTemplate, error) {
	return u.houseworkTemplateRepo.ListByFamilyID(ctx, familyID)
}

func (u *HouseworkUsecase) CreateHouseworkTemplate(ctx context.Context, template domain.HouseworkTemplate) error {
	return u.houseworkTemplateRepo.Create(ctx, template)
}

func (u *HouseworkUsecase) UpdateHouseworkTemplate(ctx context.Context, template domain.HouseworkTemplate) error {
	return u.houseworkTemplateRepo.Update(ctx, template)
}

func (u *HouseworkUsecase) DeleteHouseworkTemplate(ctx context.Context, id uint64) error {
	return u.houseworkTemplateRepo.Delete(ctx, id)
}
