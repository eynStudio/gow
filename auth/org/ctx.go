package org

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
	"github.com/eynstudio/gox/utils"
)

func init() {
	gobreak.Must(di.Reg(&OrgCtx{}))
}

type OrgCtx struct {
	*orm.Orm `di:"*"`
}

func (p *OrgCtx) Get(id gobreak.GUID) (m AuthOrg, ok bool) {
	ok = p.Orm.WhereId(id).GetJson2(&m)
	return
}

func (p *OrgCtx) All() (lst []AuthOrg, err error) {
	err = p.Orm.AllJson(&lst)
	return
}
func (p *OrgCtx) AllAsTree() (tree interface{}, err error) {
	lst := make([]AuthOrg, 0)
	if err = p.Orm.AllJson(&lst); err != nil {
		return nil, err
	}
	return utils.BuildTree(lst), nil
}

func (p *OrgCtx) Save(m *AuthOrg) gobreak.IStatus {
	err := p.Orm.SaveJson(m.Id, m)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}

func (p *OrgCtx) Del(id gobreak.GUID) gobreak.IStatus {
	err := p.Orm.DelId(&AuthOrg{}, id)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
