package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/auth/repo"
)

type OrgAgg struct {
	repo.IOrgRepo `di`
	id            GUID
	Error
}

func NewOrgAgg(id GUID) (m *OrgAgg) {
	m = &OrgAgg{id: id}
	di.Apply(m)
	return m
}
