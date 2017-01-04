package users

import (
	"encoding/json"
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
	ok = p.Orm.Where(`json->>'Mc'=? and json->>'Pwd'=?`, mc, pwd).GetJson2(&m)
	return
}
func (p *UserCtx) All() (lst []AuthUser, err error) {
	err = p.Orm.AllJson(&lst)
	return
}
func (p *UserCtx) UserCountByGroup(gid gobreak.GUID) (n int) {
	//p.Orm.Where(`json->'Groups'@>[]`)
	//	err = p.Orm.AllJson(&lst)
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

func (p *UserCtx) PageGroupUser(gid gobreak.GUID, page *filter.PageFilter) (m *db.Paging, err error) {
	lst := []UserLine{}
	s := p.Orm.From("AuthUser")
	args := db.NewAgrs(`json->'Groups' @> ?`, gid)
	if page.Search() != "" {
		str := "%" + page.Search() + "%"
		args.Append(`and (json->>'Mc' like ? or json->>'Nc' like ?)`, str, str)
	}
	m = s.Where(args.Sql, args.Args...).Select(`id ,json->>'Mc' mc,json->>'Nc' nc`).PageJson2(&lst, page)
	err = s.Err
	return
}

func (p *UserCtx) PageGroupUserSelect(gid gobreak.GUID, page *filter.PageFilter) (m *db.Paging, err error) {
	log.Println(gid, page)

	lst0 := []gobreak.GUID{gid}
	lst2, _ := json.Marshal(lst0)

	lst := []UserLine{}
	s := p.Orm.From("AuthUser")
	args := db.NewAgrs(`not json->'Groups' @> ?`, lst2)
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
	return nil
}

func (p *UserCtx) Save(m *AuthUser) gobreak.IStatus {
	err := p.Orm.SaveJson(m.Id, m)
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}

func (p *UserCtx) Del(id gobreak.GUID) gobreak.IStatus {
	err := p.Orm.DelId(&AuthUser{}, id)
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
