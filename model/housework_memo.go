package model

import (
	"gorm.io/gorm"
	"time"
)

type HouseworkMemo struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	HouseworkID   uint
	DraftedTo     uint
	DraftedToUser User   `gorm:"foreignKey:DraftedTo"`
	Message       string `json:"message" gorm:"type:text;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (m *HouseworkMemo) Create(tx *gorm.DB) error {
	return txExec("create", m, tx)
}

func (m *HouseworkMemo) Update(tx *gorm.DB) error {
	return txExec("update", m, tx)
}

func (m *HouseworkMemo) Delete(tx *gorm.DB) error {
	return txExec("delete", m, tx)
}

func (m *HouseworkMemo) Get() (*HouseworkMemo, error) {
	var res HouseworkMemo
	if err := DB.Where(m).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *HouseworkMemo) GetAll() ([]HouseworkMemo, error) {
	var res []HouseworkMemo
	if err := DB.Where(m).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
