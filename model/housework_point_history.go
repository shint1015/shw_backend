package model

import (
	"gorm.io/gorm"
	"time"
)

type HouseworkPointHistory struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	UserID         uint   `json:"user_id" gorm:"foreignKey:UserID"`
	TargetUser     User   `json:"target_user_id" gorm:"foreignKey:UserID"`
	Detail         string `json:"detail" gorm:"type:text;not null"`
	Point          int    `json:"point" gorm:"type:int;not null"`
	AggregatedFlag bool   `json:"aggregated_flag" gorm:"type:bool;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (m *HouseworkPointHistory) Create(tx *gorm.DB) error {
	return txExec("create", m, tx)
}

func (m *HouseworkPointHistory) Update(tx *gorm.DB) error {
	return txExec("update", m, tx)
}

func (m *HouseworkPointHistory) Delete(tx *gorm.DB) error {
	return txExec("delete", m, tx)
}

func (m *HouseworkPointHistory) Get() (*HouseworkPointHistory, error) {
	var res HouseworkPointHistory
	if err := DB.Where(m).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *HouseworkPointHistory) GetAll(whereQuery [][]interface{}) ([]HouseworkPointHistory, error) {
	var res []HouseworkPointHistory
	query := DB.Where(m)
	for _, v := range whereQuery {
		if v == nil || len(v) == 0 {
			continue
		}
		query = query.Where(v[0], v[1:]...)
	}
	if err := query.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
