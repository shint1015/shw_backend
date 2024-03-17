package model

import (
	"gorm.io/gorm"
	"time"
)

type FamilyRole struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	FamilyID  uint
	Family    Family `json:"family_id" gorm:"foreignKey:FamilyID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *FamilyRole) Create(tx *gorm.DB) error {
	return txExec("create", m, tx)
}

func (m *FamilyRole) Update(tx *gorm.DB) error {
	return txExec("update", m, tx)
}

func (m *FamilyRole) Delete(tx *gorm.DB) error {
	return txExec("delete", m, tx)
}

func (m *FamilyRole) Get() (*FamilyRole, error) {
	var res FamilyRole
	if err := DB.Where(m).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *FamilyRole) GetAll() ([]FamilyRole, error) {
	var res []FamilyRole
	if err := DB.Where(m).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
