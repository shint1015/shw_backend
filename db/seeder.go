package db

import (
	"shwgrpc/model"
	"shwgrpc/utils"
)

type Seeder struct{}

func DoSeed() {
	s := Seeder{}
	s.user()
	s.familyRole()
	s.family()
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
	families = append(families, model.Family{Name: "Smith", PointPerWorkTime: "1", OwnerUserID: 1})
	families = append(families, model.Family{Name: "Johnson", PointPerWorkTime: "5", OwnerUserID: 2})
	families = append(families, model.Family{Name: "Willams", PointPerWorkTime: "10", OwnerUserID: 3})

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
	familyRoles = append(familyRoles, model.FamilyRole{Name: "father", 1})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "mother", 1})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "brother", 1})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "father", 2})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "mother", 2})
	familyRoles = append(familyRoles, model.FamilyRole{Name: "father", 3})
	tx := model.DB.Begin()
	for _, familyRole := range familyRoles {
		if err := familyRole.Create(tx); err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit();
}

func (s *Seeder) housework() {
	var houseworks []model.Housework
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "洗濯", Detail: "家族の洋服", Status: "plan", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "掃除", Detail: "リビング掃除", Status: "plan", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 11:30:00 JST"), EndedAt: utils.Str2Time("2024-03-08 12:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "食事作り", Detail: "カレーライス", Status: "plan", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 12:30:00 JST"), EndedAt: utils.Str2Time("2024-03-08 13:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "洗濯", Detail: "子供の洋服", Status: "doing", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-08 16:30:00 JST"), EndedAt: utils.Str2Time("2024-03-08 17:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "掃除", Detail: "トイレ掃除", Status: "doing", WorkTo: 4, StartedAt: utils.Str2Time("2024-03-08 18:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 19:00:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "食事作り", Detail: "ナポリタン", Status: "doing", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 20:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 21:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "洗濯", Detail: "家族の洋服", Status: "plan", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-09 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-09 11:30:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "掃除", Detail: "キッチン掃除", Status: "plan", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-09 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-09 12:00:00 JST")})
	houseworks = append(houseworks, model.Housework{FamilyID: 1, Title: "食事作り", Detail: "ハンバーグ", Status: "plan", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-09 13:00:00 JST"), EndedAt: utils.Str2Time("2024-03-09 13:30:00 JST")})
}

func (s *Seeder) houseworkMemo() {
	var houseworkMemos []model.HouseworkMemo
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 1, DraftedTo: , Message: });
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 1, DraftedTo: , Message: });
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 3, DraftedTo: , Message: });
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 3, DraftedTo: , Message: });
	houseworkMemos = append(houseworkMemos, model.HouseworkMemo{HouseworkID: 3, DraftedTo: , Message: });
}

func (s *Seeder) houseworkTemplate() {

}

func (s *Seeder) houseworkPoint() {

}
