package db

import (
	"shwgrpc/model"
	"shwgrpc/utils"
)

type Seeder struct{}

func DoSeed() {
	s := Seeder{}
	s.user()
	s.family()
	s.familyRole()
	s.housework()
	s.houseworkMemo()
	s.houseworkTemplate()
	s.houseworkPoint()
}

func (s *Seeder) user() {
	var users []model.User
	users = append(users, model.User{Name: "Tom", Password: "password", Email: "tom@samle.com"})
	users = append(users, model.User{Name: "Brown", Password: "password", Email: "brown@samle.com"})
	users = append(users, model.User{Name: "Peter", Password: "password", Email: "peter@samle.com"})
	users = append(users, model.User{Name: "John", Password: "password", Email: "john@samle.com"})
	users = append(users, model.User{Name: "Mary", Password: "password", Email: "mary@samle.com"})
	users = append(users, model.User{Name: "Alice", Password: "password", Email: "alice@samle.com"})
	tx := model.DB.Begin()
	for _, user := range users {
		if err := user.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}

func (s *Seeder) family() {
	var families []model.Family
	var ownerOne uint = 1
	var ownerTwo uint = 2
	var ownerThree uint = 3
	families = append(families, model.Family{Name: "Smith", PointPerWorkTime: "1", OwnerUserID: &ownerOne})
	families = append(families, model.Family{Name: "Johnson", PointPerWorkTime: "5", OwnerUserID: &ownerTwo})
	families = append(families, model.Family{Name: "Willams", PointPerWorkTime: "10", OwnerUserID: &ownerThree})

	tx := model.DB.Begin()
	for _, family := range families {
		if err := family.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}

func (s *Seeder) familyRole() {
	var familyRoles []model.FamilyRole
	familyRoles = append(familyRoles, model.FamilyRole{Name: "father", FamilyID: 1})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "mother", FamilyID: 1})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "brother", FamilyID: 1})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "father", FamilyID: 2})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "mother", FamilyID: 2})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "father", FamilyID: 3})
	tx := model.DB.Begin()
	for _, familyRole := range familyRoles {
		if err := familyRole.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}

func (s *Seeder) housework() {
	var houseworks []model.Housework
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "洗濯", Detail: "家族の洋服", Status: "done", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "掃除", Detail: "リビング掃除", Status: "done", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 11:30:00 JST"), EndedAt: utils.Str2Time("2024-03-08 12:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "食事作り", Detail: "カレーライス", Status: "done", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 12:30:00 JST"), EndedAt: utils.Str2Time("2024-03-08 13:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "洗濯", Detail: "子供の洋服", Status: "doing", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-08 16:30:00 JST"), EndedAt: utils.Str2Time("2024-03-08 17:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "掃除", Detail: "トイレ掃除", Status: "doing", WorkTo: 4, StartedAt: utils.Str2Time("2024-03-08 18:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 19:00:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "食事作り", Detail: "ナポリタン", Status: "doing", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 20:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 21:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "洗濯", Detail: "家族の洋服", Status: "plan", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-09 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-09 11:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "掃除", Detail: "キッチン掃除", Status: "plan", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-09 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-09 12:00:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "食事作り", Detail: "ハンバーグ", Status: "plan", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-09 13:00:00 JST"), EndedAt: utils.Str2Time("2024-03-09 13:30:00 JST")})

	tx := model.DB.Begin()
	for _, housework := range houseworks {
		if err := housework.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}

func (s *Seeder) houseworkMemo() {
	var houseworkMemos []model.HouseworkMemo
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 1, SendFrom: 5, Message: "ありがとう"})
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 1, SendFrom: 4, Message: "ありがとう"})
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 3, SendFrom: 5, Message: "美味しかった、ありがとう"})
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 3, SendFrom: 4, Message: "美味しかった、ありがとう"})
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 3, SendFrom: 1, Message: "どういたしまして"})

	tx := model.DB.Begin()
	for _, houseworkMemo := range houseworkMemos {
		if err := houseworkMemo.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}

func (s *Seeder) houseworkTemplate() {
	var houseworkTemplates []model.HouseworkTemplate
	houseworkTemplates = append(houseworkTemplates, model.HouseworkTemplate{FamilyID: 1, Title: "洗濯", Detail: "家族の洋服"})
	houseworkTemplates = append(houseworkTemplates, model.HouseworkTemplate{FamilyID: 1, Title: "掃除", Detail: "リビング掃除"})

	tx := model.DB.Begin()
	for _, houseworkTemplate := range houseworkTemplates {
		if err := houseworkTemplate.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()

}

func (s *Seeder) houseworkPoint() {
	var houseworkPoints []model.HouseworkPoint
	houseworkPoints = append(houseworkPoints, model.HouseworkPoint{UserID: 1, Point: 10})
	houseworkPoints = append(houseworkPoints, model.HouseworkPoint{UserID: 4, Point: 0})
	houseworkPoints = append(houseworkPoints, model.HouseworkPoint{UserID: 5, Point: 20})

	tx := model.DB.Begin()
	for _, houseworkPoint := range houseworkPoints {
		if err := houseworkPoint.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()

}

func (s *Seeder) houseworkPointHistory() {
	var houseworkIdOne uint = 1
	var houseworkIdTwo uint = 2
	var houseworkIdThree uint = 3
	var houseworkPointHistories []model.HouseworkPointHistory
	houseworkPointHistories = append(houseworkPointHistories, model.HouseworkPointHistory{UserID: 1, HouseworkID: &houseworkIdOne, Detail: "ポイント付与", Point: 10, AggregatedFlag: true})
	houseworkPointHistories = append(houseworkPointHistories, model.HouseworkPointHistory{UserID: 5, HouseworkID: &houseworkIdTwo, Detail: "ポイント付与", Point: 10, AggregatedFlag: true})
	houseworkPointHistories = append(houseworkPointHistories, model.HouseworkPointHistory{UserID: 5, HouseworkID: &houseworkIdThree, Detail: "ポイント付与", Point: 10, AggregatedFlag: true})

	tx := model.DB.Begin()
	for _, houseworkPointHistory := range houseworkPointHistories {
		if err := houseworkPointHistory.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()

}
