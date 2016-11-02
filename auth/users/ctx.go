package users

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

func init() {
	log.Println(di.Reg(&UserCtx{}))
}

type UserCtx struct {
	*orm.Orm `di:"*"`
}

func (p *UserCtx) Get(id gobreak.GUID) (m AuthUser, ok bool) {
	ok = p.Orm.WhereId(id).GetJson2(&m)
	return
}

func (p *UserCtx) All() (lst []AuthUser, err error) {
	err = p.Orm.AllJson(&lst)
	return
}

func (p *UserCtx) Save(m *AuthUser) gobreak.IStatus {
	err := p.Orm.SaveJson(m.Id, m)
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}

func (p *UserCtx) Del(id gobreak.GUID) gobreak.IStatus {
	err := p.Orm.DelId(&AuthUser{}, id)
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
