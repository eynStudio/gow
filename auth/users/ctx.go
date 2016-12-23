package users

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/db"
	"github.com/eynstudio/gobreak/db/filter"
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
func (p *UserCtx) GetByMcPwd(mc, pwd string) (m AuthUser, ok bool) {
	ok = p.Orm.Where(`json->>'mc'=? and json->>'pwd'=?`, mc, pwd).GetJson2(&m)
	return
}
func (p *UserCtx) All() (lst []AuthUser, err error) {
	err = p.Orm.AllJson(&lst)
	return
}

func (p *UserCtx) PageUser(page *filter.PageFilter) (m *db.Paging, err error) {

	lst := []UserLine{}

	s := p.Orm.From("AuthUser")
	if page.Search() != "" {
		str := "%" + page.Search() + "%"
		s.Where(`json->>'Mc' like ? or json->>'Nc' like ?`, str, str)
	}
	m = s.Select(`id ,json->>'Mc' mc,json->>'Nc' nc`).Page2(&lst, page)
	err = s.Err
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
