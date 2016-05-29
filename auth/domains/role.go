package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/auth/repo"
)

type RoleAgg struct {
	repo.IRoleRepo `di`
	id             GUID
	Error
}

func NewRoleAgg(id GUID) (m *RoleAgg) {
	m = &RoleAgg{id: id}
	di.Apply(m)
	return m
}
