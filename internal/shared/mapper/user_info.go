package mapper

import (
	"shwgrpc/model"
)

func MapUserInfoAs[T any](u model.User, build func(id uint64, name string) T) T {
	return build(uint64(u.ID), u.Name)
}
