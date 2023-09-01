package model

import (
	"gorm.io/gorm"
	"time"
)

type HouseworkTemplate struct {
	ID        uint `json:"id" gorm:"primary_key"`
	FamilyID  uint
	Name      string `json:"name" gorm:"unique;type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *HouseworkTemplate) Create(tx *gorm.DB) error {
	return txExec("create", m, tx)
}

func (m *HouseworkTemplate) Update(tx *gorm.DB) error {
	return txExec("update", m, tx)
}

func (m *HouseworkTemplate) Delete(tx *gorm.DB) error {
	return txExec("delete", m, tx)
}

func (m *HouseworkTemplate) Get() (*HouseworkTemplate, error) {
	var res HouseworkTemplate
	if err := DB.Where(m).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *HouseworkTemplate) GetAll() ([]HouseworkTemplate, error) {
	var res []HouseworkTemplate
	if err := DB.Where(m).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
