package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/cms/repo"
)

type InfoAgg struct {
	repo.IInfoRepo `di`
	id             GUID
	Error
}

func NewInfoAgg(id GUID) (m *InfoAgg) {
	m = &InfoAgg{id: id}
	di.Apply(m)
	return m
}
