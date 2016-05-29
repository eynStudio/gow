package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/auth/repo"
)

type UserAgg struct {
	r  repo.IUserRepo `di`
	id GUID
	Error
}

func NewUserAgg(id GUID) (m *UserAgg) {
	m = &UserAgg{id: id}
	di.Apply(m)
	return m
}

func (p *UserAgg) UpdateNc(nc string) {
	p.r.UpdateNc(p.id, nc)
}

//Get 昵称+头像
func (p *UserAgg) GetNcImg() (nc string, img string) {
	u, _ := p.r.GetById(p.id)
	if u.Nc == "" {
		return u.Mc, u.Img
	}
	return u.Nc, u.Img
}
