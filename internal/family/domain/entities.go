package domain

import (
	houseworkdomain "shwgrpc/internal/housework/domain"
)

type Family struct {
	ID    *uint64
	Users []houseworkdomain.UserInfo
	Name  string
}

type FamilyRole struct {
	ID *uint64
	Name string 
}