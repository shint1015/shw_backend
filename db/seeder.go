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

}

// TODO: familyroleのほうがいいかも
func (s *Seeder) userRole() {

}

func (s *Seeder) family() {

}

func (s *Seeder) housework() {

}

func (s *Seeder) houseworkMemo() {

}

func (s *Seeder) houseworkTemplate() {

}

func (Seeder *Seeder) houseworkPoint() {

}
