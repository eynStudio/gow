package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
)

type InfoAgg struct {
	id GUID
	Error
}

func NewInfoAgg(id GUID) (m *InfoAgg) {
	m = &InfoAgg{id: id}
	di.Apply(m)
	return m
}
