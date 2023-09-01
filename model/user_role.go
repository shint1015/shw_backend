package model

import (
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"unique;type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *UserRole) Create(tx *gorm.DB) error {
	return txExec("create", m, tx)
}

func (m *UserRole) Update(tx *gorm.DB) error {
	return txExec("update", m, tx)
}

func (m *UserRole) Delete(tx *gorm.DB) error {
	return txExec("delete", m, tx)
}

func (m *UserRole) Get() (*UserRole, error) {
	var res UserRole
	if err := DB.Where(m).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *UserRole) GetAll() ([]UserRole, error) {
	var res []UserRole
	if err := DB.Where(m).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
