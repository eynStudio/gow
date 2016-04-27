package auth

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"

	"github.com/eynstudio/gow/auth/role"
)

type RoleRepo struct {
	MgoRepo
}

func NewRoleRepo() *RoleRepo {
	return &RoleRepo{NewMgoRepo("AuthRole", func() T { return role.NewRole() })}
}
