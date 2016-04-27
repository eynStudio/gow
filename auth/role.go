package auth

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	"gopkg.in/mgo.v2/bson"

	"github.com/eynstudio/gow/auth/role"
)

type RoleRepo struct {
	MgoRepo
}

func NewRoleRepo() *RoleRepo {
	return &RoleRepo{NewMgoRepo("AuthRole", func() T { return role.NewRole() })}
}

func (p *RoleRepo) GetRoleIdByUri(uri string) GUID {
	var r role.Role
	p.GetQAs(&r, bson.M{"Uri": uri})
	return r.Id
}
