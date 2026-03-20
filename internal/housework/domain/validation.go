package domain

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

const maxHouseworkTextLength = 255

var (
	ErrHouseworkIDRequired        = errors.New("housework id is required")
	ErrHouseworkFamilyIDRequired  = errors.New("family_id is required")
	ErrHouseworkTitleRequired     = errors.New("title is required")
	ErrHouseworkWorkUserIDRequired = errors.New("work_user.id is required")
	ErrHouseworkStatusInvalid     = errors.New("status is invalid")
	ErrHouseworkTimeRangeInvalid  = errors.New("ended_at must be same or after started_at")
)

func NewHouseworkForCreate(h Housework) (Housework, error) {
	h = h.normalizeForCreate()
	if err := h.ValidateForCreate(); err != nil {
		return Housework{}, err
	}
	return h, nil
}

func NewHouseworkForUpdate(h Housework) (Housework, error) {
	if err := h.ValidateForUpdate(); err != nil {
		return Housework{}, err
	}
	return h, nil
}

func (h Housework) ValidateForCreate() error {
	if h.FamilyID == 0 {
		return ErrHouseworkFamilyIDRequired
	}
	if err := h.validateCommonFields(); err != nil {
		return err
	}
	if h.Status == "" || HouseworkStatusToID(h.Status) == HouseworkStatusIDUnspecified {
		return fmt.Errorf("%w: %s", ErrHouseworkStatusInvalid, h.Status)
	}
	if h.StatusID != HouseworkStatusToID(h.Status) {
		return fmt.Errorf("%w: status=%s status_id=%d", ErrHouseworkStatusInvalid, h.Status, h.StatusID)
	}
	return nil
}

func (h Housework) ValidateForUpdate() error {
	if h.ID == 0 {
		return ErrHouseworkIDRequired
	}
	if err := h.validateCommonFields(); err != nil {
		return err
	}
	return nil
}

func (h Housework) validateCommonFields() error {
	if strings.TrimSpace(h.Title) == "" {
		return ErrHouseworkTitleRequired
	}
	if utf8.RuneCountInString(h.Title) > maxHouseworkTextLength {
		return fmt.Errorf("title must be %d characters or less", maxHouseworkTextLength)
	}
	if utf8.RuneCountInString(h.Detail) > maxHouseworkTextLength {
		return fmt.Errorf("detail must be %d characters or less", maxHouseworkTextLength)
	}
	if h.WorkUser.ID == 0 {
		return ErrHouseworkWorkUserIDRequired
	}
	if !h.StartedAt.IsZero() && !h.EndedAt.IsZero() && h.EndedAt.Before(h.StartedAt) {
		return ErrHouseworkTimeRangeInvalid
	}
	return nil
}

func (h Housework) normalizeForCreate() Housework {
	if h.Status == "" && h.StatusID != 0 {
		h.Status = HouseworkStatusFromID(h.StatusID)
	}
	if h.Status == "" {
		h.Status = HouseworkStatusPlan
	}
	h.StatusID = HouseworkStatusToID(h.Status)
	return h
}
