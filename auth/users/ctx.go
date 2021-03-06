package users

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/db"
	"github.com/eynstudio/gobreak/db/filter"
	"github.com/eynstudio/gobreak/orm"
)

var Ctx = &UserCtx{}

type UserCtx struct {
}

func (p *UserCtx) orm() *orm.Orm { return orm.GetOrmByName("auth") }

func (p *UserCtx) Get(id gobreak.GUID) (m AuthUser, ok bool) {
	ok = p.orm().WhereId(id).GetJson2(&m)
	return
}
func (p *UserCtx) GetByMcPwd(mc, pwd string) (m AuthUser, ok bool) {
	ok = p.orm().Where(`json->>'Mc'=? and json->>'Pwd'=?`, mc, pwd).GetJson2(&m)
	return
}
func (p *UserCtx) All() (lst []AuthUser, err error) {
	err = p.orm().AllJson(&lst)
	return
}
func (p *UserCtx) UserCountByGroup(gid gobreak.GUID) (n int) {
	//p.Orm.Where(`json->'Groups'@>[]`)
	//	err = p.Orm.AllJson(&lst)
	return
}
func (p *UserCtx) PageUser(page *filter.PageFilter) (m *db.Paging, err error) {
	lst := []UserLine{}
	s := p.orm().From("AuthUser")
	if page.Search() != "" {
		str := "%" + page.Search() + "%"
		s.Where(`json->>'Mc' like ? or json->>'Nc' like ?`, str, str)
	}
	m = s.Select(`id ,json->>'Mc' mc,json->>'Nc' nc`).Page2(&lst, page)
	err = s.Err
	return
}

func (p *UserCtx) PageGroupUser(gid gobreak.GUID, page *filter.PageFilter) (m *db.Paging, err error) {
	return p.pageGroupUser(gid, page, true)
}

func (p *UserCtx) PageGroupUserSelect(gid gobreak.GUID, page *filter.PageFilter) (m *db.Paging, err error) {
	return p.pageGroupUser(gid, page, false)
}

func (p *UserCtx) pageGroupUser(gid gobreak.GUID, page *filter.PageFilter, in bool) (m *db.Paging, err error) {
	lst := []UserLine{}
	s := p.orm().From("AuthUser")
	sql := gobreak.IfThenStr(in, "", "not ") + `json->'Groups' @> ?`
	args := db.NewAgrs(sql, gid.Json())
	if page.Search() != "" {
		str := "%" + page.Search() + "%"
		args.Append(`and (json->>'Mc' like ? or json->>'Nc' like ?)`, str, str)
	}
	m = s.Where(args.Sql, args.Args...).Select(`id, json->>'Mc' mc,json->>'Nc' nc`).Page2(&lst, page)
	err = s.Err
	return
}

func (p *UserCtx) AddUserGroup(uid, gid gobreak.GUID) error {
	log.Println(uid, gid)
	if u, ok := p.Get(uid); ok {
		u.AddGroup(gid)
		return p.orm().SaveJson(u.Id, u)
	}
	return nil
}

func (p *UserCtx) Save(m *AuthUser) gobreak.IStatus {
	err := p.orm().SaveJson(m.Id, m)
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}

func (p *UserCtx) Del(id gobreak.GUID) gobreak.IStatus {
	err := p.orm().DelId(&AuthUser{}, id)
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
