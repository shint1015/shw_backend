package db

type Seeder struct{}

func DoSeed() {
	s := Seeder{}
	s.user()
	s.userRole()
	s.family()
	s.housework()
	s.houseworkMemo()
	s.houseworkTemplate()
	s.houseworkPoint()
}

func (s *Seeder) user() {
	var users []User
	users = append(users, User{Name: "Tom", Password: "password", Email: "tom@samle.com"})
	users = append(users, User{Name: "Brown", Password: "password", Email: "brown@samle.com"})
	users = append(users, User{Name: "Peter", Password: "password", Email: "peter@samle.com"})
	users = append(users, User{Name: "John", Password: "password", Email: "john@samle.com"})
	users = append(users, User{Name: "Mary", Password: "password", Email: "mary@samle.com"})
	users = append(users, User{Name: "Alice", Password: "password", Email: "alice@samle.com"})
	for _, user := range users {
		user.create();
	}
}

// TODO: familyroleのほうがいいかも
func (s *Seeder) userRole() {

}

func (s *Seeder) family() {
	var families []Family
	families = append(families, Family{Name: "Smith", PointPerWorkTime: 1, OwnerUserID: 1})
	families = append(families, Family{Name: "Johnson", PointPerWorkTime: 5, OwnerUserID: 2})
	families = append(families, Family{Name: "Willams", PointPerWorkTime: 10, OwnerUserID: 3})
	for _, family := range 
}

func (s *Seeder) housework() {
	var houseworks []Housework
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "洗濯", Detail: "家族の洋服", Status: "plan", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "掃除", Detail: "リビング掃除", Status: "plan", WorkTo: 5,  StartedAt: utils.Str2Time("2024-03-08 11:30:00 JST"), EndedAt: utils.Str2Time("2024-03-08 12:30:00 JST")})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "食事作り", Detail: "カレーライス", Status: "plan", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 12:30:00 JST"), EndedAt: utils.Str2Time("2024-03-08 13:30:00 JST")})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "洗濯", Detail: "子供の洋服", Status: "doing", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "掃除", Detail: "トイレ掃除", Status: "doing", WorkTo: 4, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "食事作り", Detail: "ナポリタン", Status: "doing", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "洗濯", Detail: "家族の洋服", Status: "plan", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "掃除", Detail: "キッチン掃除", Status: "plan", WorkTo: 1, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "食事作り", Detail: "ハンバーグ", Status: "plan", WorkTo: 5, StartedAt: utils.Str2Time("2024-03-08 11:00:00 JST"), EndedAt: utils.Str2Time("2024-03-08 11:30:00 JST")})
}

func (s *Seeder) houseworkMemo() {

}

func (s *Seeder) houseworkTemplate() {

}

func (s *Seeder) houseworkPoint() {

}
