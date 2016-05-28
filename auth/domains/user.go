package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/auth/repo"
)

type UserAgg struct {
	repo.IUserRepo `di`
	id             GUID
	Error
}

func NewUserAgg(id GUID) (m *UserAgg) {
	m = &UserAgg{id: id}
	di.Apply(m)
	return m
}
