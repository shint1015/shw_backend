package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func init() {
	USER, _ := os.LookupEnv("DATABASE_USER")
	PASS, _ := os.LookupEnv("DATABASE_PASSWORD")
	PROTOCOL, _ := os.LookupEnv("DATABASE_PROTOCOL")
	DBNAME, _ := os.LookupEnv("DATABASE_NAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	DB, err = gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to DB %s: %v", CONNECT, err)
	}

	if err = DB.AutoMigrate(&UserRole{}, &Family{}, &User{}, &Housework{}, &HouseworkMemo{}, &HouseworkTemplate{}, &HouseworkPoint{}, &HouseworkPointHistory{}); err != nil {
		log.Fatalf("Error migrating DB %s: %v", CONNECT, err)
	}
}

func txExec[
	modelData *User |
		*UserRole |
		*Family |
		*Housework |
		*HouseworkMemo |
		*HouseworkTemplate |
		*HouseworkPoint |
		*HouseworkPointHistory](
	queryType string,
	m modelData,
	tx *gorm.DB,
) error {
	existTx := true
	if tx == nil {
		tx = DB.Begin()
		existTx = false
	}
	var res *gorm.DB
	switch queryType {
	case "create":
		res = tx.Create(&m)
	case "update":
		res = tx.Save(&m)
	case "delete":
		res = tx.Delete(&m)
	}

	if res.Error != nil {
		if !existTx {
			tx.Rollback()
		}
		return res.Error
	}
	if !existTx {
		tx.Commit()
	}
	return nil
}
