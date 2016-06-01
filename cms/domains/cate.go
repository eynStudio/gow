package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/cms/repo"
)

type CateAgg struct {
	repo.ICateRepo `di`
	id             GUID
	Error
}

func NewCateAgg(id GUID) (m *CateAgg) {
	m = &CateAgg{id: id}
	di.Apply(m)
	return m
}
