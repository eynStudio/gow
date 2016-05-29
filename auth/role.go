package auth

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gow/auth/repo"

	. "github.com/eynstudio/gobreak/db/mgo"
	. "github.com/eynstudio/gow/auth/models"
	"gopkg.in/mgo.v2/bson"
)

var _ repo.IRoleRepo = new(RoleRepo)

type RoleRepo struct {
	MgoRepo
}

func NewRoleRepo() *RoleRepo {
	return &RoleRepo{NewMgoRepo("AuthRole", func() T { return NewRole() })}
}

func (p *RoleRepo) GetRoleIdByUri(uri string) GUID {
	var r Role
	p.GetQAs(&r, bson.M{"Uri": uri})
	return r.Id
}
