package repo

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
)

type IRoleRepo interface {
	MgoRepo
	GetRoleIdByUri(uri string) GUID
}
