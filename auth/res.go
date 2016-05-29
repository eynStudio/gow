package auth

import (
	. "github.com/eynstudio/gow/auth/models"
	"github.com/eynstudio/gow/auth/repo"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/db"
	. "github.com/eynstudio/gobreak/db/mgo"
)

var _ repo.IResRepo = new(ResRepo)

type ResRepo struct {
	MgoRepo
}

func NewResRepo() *ResRepo {
	return &ResRepo{NewMgoRepo("AuthRes", func() T { return &Res{} })}
}

func (p *ResRepo) GetById(id GUID) (m *Res, err error) {
	m = new(Res)
	p.GetAs(id, &m)
	if m.Id == "" {
		err = db.DbNotFound
	}
	return
}
