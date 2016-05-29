package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/auth/models"
	"github.com/eynstudio/gow/auth/repo"
)

type ResAgg struct {
	repo.IResRepo `di`
	id            GUID
	Error
	root *models.Res
}

func NewResAgg(id GUID) (m *ResAgg) {
	m = &ResAgg{id: id}
	di.Apply(m)
	return m
}

func (p *ResAgg) LoadRoot() {
	p.root, p.Err = p.GetById(p.id)
}

func (p *ResAgg) Save(m models.Res) {

}

func (p *ResAgg) Del() {

}
