package model

import (
	"gorm.io/gorm"
	"time"
)

type HouseworkPoint struct {
	ID        uint `json:"id" gorm:"primary_key"`
	UserID    uint
	User      User `json:"user_id" gorm:"foreignKey:UserID"`
	Point     int  `json:"point" gorm:"type:int;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *HouseworkPoint) Create(tx *gorm.DB) error {
	return txExec("create", m, tx)
}

func (m *HouseworkPoint) Update(tx *gorm.DB) error {
	return txExec("update", m, tx)
}

func (m *HouseworkPoint) Delete(tx *gorm.DB) error {
	return txExec("delete", m, tx)
}

func (m *HouseworkPoint) Get() (*HouseworkPoint, error) {
	var res HouseworkPoint
	if err := DB.Where(m).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *HouseworkPoint) GetAll() ([]HouseworkPoint, error) {
	var res []HouseworkPoint
	if err := DB.Where(m).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
