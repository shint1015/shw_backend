package model

import (
	"gorm.io/gorm"
	"time"
)

type PointPerWorkTimeNum int

type Family struct {
	ID                 uint                `json:"id" gorm:"primaryKey"`
	Users              *[]User             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name               string              `json:"name" gorm:"unique;type:varchar(255);not null"`
	PointPerWorkTime   string              `json:"point_per_work_time" gorm:"type:enum('1', '5', '10', '15', '30', '60'); default: '10'; not null"`
	Housework          []Housework         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HouseworkTemplates []HouseworkTemplate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OwnerUserID        uint
	OwnerUser          *User `gorm:"foreignKey:OwnerUserID"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func (m *Family) Create(tx *gorm.DB) error {
	return txExec("create", m, tx)
}

func (m *Family) Update(tx *gorm.DB) error {
	return txExec("update", m, tx)
}

func (m *Family) Delete(tx *gorm.DB) error {
	return txExec("delete", m, tx)
}

func (m *Family) Get() (*Family, error) {
	var res Family
	if err := DB.Where(m).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *Family) GetAll() ([]Family, error) {
	var res []Family
	if err := DB.Where(m).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
