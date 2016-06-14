package domains

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/auth/repo"
)

type UserAgg struct {
	id   GUID
	Reop repo.IUserRepo `di`
	Error
}

func NewUserAgg(id GUID) (m *UserAgg) {
	m = &UserAgg{id: id}
	di.Apply(m)
	return m
}

func (p *UserAgg) UpdateNc(nc string)   { p.Reop.UpdateNc(p.id, nc) }
func (p *UserAgg) UpdatePwd(pwd string) { p.Reop.UpdatePwd(p.id, pwd) }
func (p *UserAgg) UpdateImg(img string) { p.Reop.UpdateImg(p.id, img) }
func (p *UserAgg) Lock()                { p.Reop.UpdateLock(p.id, true) }
func (p *UserAgg) Unlock()              { p.Reop.UpdateLock(p.id, false) }
func (p *UserAgg) CheckThenUpdatePwd(pwd0, pwd1 string) {
	has, err := p.Reop.CheckPwd(p.id, pwd0)
	if err != nil {
		p.Err = err
		return
	}
	if !has {
		p.SetErr("原密码错误")
		return
	}
	p.Reop.UpdatePwd(p.id, pwd1)
}

//Get 昵称+头像
func (p *UserAgg) GetNcImg() (nc string, img string) {
	u, _ := p.Reop.GetById(p.id)
	if u.Nc == "" {
		return u.Mc, u.Img
	}
	return u.Nc, u.Img
}
