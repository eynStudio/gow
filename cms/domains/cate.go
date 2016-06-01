package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
)

type CateAgg struct {
	id GUID
	Error
}

func NewCateAgg(id GUID) (m *CateAgg) {
	m = &CateAgg{id: id}
	di.Apply(m)
	return m
}
