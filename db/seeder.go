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
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "洗濯", Detail: , Status: , WorkTo: ,})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "掃除", Detail: , Status: , WorkTo: ,})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "食事作り", Detail: , Status: , WorkTo: ,})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "洗濯", Detail: , Status: , WorkTo: ,})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "掃除", Detail: , Status: , WorkTo: ,})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "食事作り", Detail: , Status: , WorkTo: ,})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "洗濯", Detail: , Status: , WorkTo: ,})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "掃除", Detail: , Status: , WorkTo: ,})
	houseworks = append(houseworks, Housework{FamilyID: 1, Title: "食事作り", Detail: , Status: , WorkTo: ,})
}

func (s *Seeder) houseworkMemo() {

}

func (s *Seeder) houseworkTemplate() {

}

func (s *Seeder) houseworkPoint() {

}
