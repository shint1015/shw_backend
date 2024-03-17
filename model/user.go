package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                   uint   `gorm:"primaryKey"`
	Name                 string `json:"name" gorm:"type:varchar(255);not null"`
	Password             string `json:"password" gorm:"type:varchar(255);not null"`
	Email                string `json:"email" gorm:"type:varchar(255);not null"`
	FamilyID             *uint
	RoleID               *uint
	Family               *Family
	Role                 *FamilyRole     `gorm:"foreignKey:RoleID"`
	HouseworkPoint       *HouseworkPoint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsFamilyVerified     bool
	FamilyVerifyExpireAt *time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (u *User) Create(tx *gorm.DB) error {
	return txExec("create", u, tx)
}

func (u *User) Update(tx *gorm.DB) error {
	return txExec("update", u, tx)
}

func (u *User) Delete(tx *gorm.DB) error {
	return txExec("delete", u, tx)
}

func (u *User) Get() (*User, error) {
	var res User
	if err := DB.Where(u).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *User) GetAll() ([]User, error) {
	var res []User
	if err := DB.Where(u).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (u *User) GetUsersByFamilyID() ([]User, error) {
	var res []User
	if err := DB.Where("family_id = ?", u.FamilyID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
