package port

type CreateFamilyInput struct {
	Name string
}

type UpdateFamilyInput struct {
	ID   uint64
	Name string
}

type CreateRoleInput struct {
	Name string
	FamilyID uint64
}

type UpdateRoleInput struct {
	ID uint64
	Name string
	FamilyID uint64
}
