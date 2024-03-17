package model

import (
	"gorm.io/gorm"
	"time"
)

type Housework struct {
	ID           uint `json:"id" gorm:"primary_key"`
	FamilyID     uint
	Family       Family `json:"family_id" gorm:"foreignKey:FamilyID"`
	Title        string `json:"title" gorm:"type:varchar(255);not null"`
	Detail       string `json:"detail" gorm:"type:varchar(255);not null"`
	Status       string `json:"status" gorm:"type:varchar(255);not null"`
	Memo         *[]HouseworkMemo
	WorkTo       uint
	WorkToUser   User `json:"work_to_id" gorm:"foreignKey:WorkTo"`
	WorkTimeNum  uint
	WorkTimeType string
	IsPointAdded bool
	StartedAt    time.Time
	EndedAt      time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (m *Housework) Create(tx *gorm.DB) error {
	return txExec("create", m, tx)
}

func (m *Housework) Update(tx *gorm.DB) error {
	return txExec("update", m, tx)
}

func (m *Housework) Delete(tx *gorm.DB) error {
	return txExec("delete", m, tx)
}

func (m *Housework) Get() (*Housework, error) {
	var res Housework
	if err := DB.Where(m).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *Housework) GetAll() ([]Housework, error) {
	var res []Housework
	if err := DB.Where(m).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
