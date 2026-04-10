package port

type CreateFamilyInput struct {
	Name string
}

type UpdateFamilyInput struct {
	ID   uint64
	Name string
}

type AddFamilyMemberInput struct {
	Name string
	Email *string
	FamilyID uint64
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
