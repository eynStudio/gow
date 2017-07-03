package org

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/utils"
)

var Ctx = &OrgCtx{}

type OrgCtx struct {
}

func (p *OrgCtx) orm() *orm.Orm { return orm.GetOrmByName("auth") }

func (p *OrgCtx) Get(id gobreak.GUID) (m AuthOrg, ok bool) {
	ok = p.orm().WhereId(id).GetJson2(&m)
	return
}

func (p *OrgCtx) All() (lst []AuthOrg, err error) {
	err = p.orm().AllJson(&lst)
	return
}
func (p *OrgCtx) AllAsTree() (tree interface{}, err error) {
	lst := make([]AuthOrg, 0)
	if err = p.orm().AllJson(&lst); err != nil {
		return nil, err
	}
	return utils.BuildTree(lst), nil
}

func (p *OrgCtx) Save(m *AuthOrg) gobreak.IStatus {
	err := p.orm().SaveJson(m.Id, m)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}

func (p *OrgCtx) Del(id gobreak.GUID) gobreak.IStatus {
	err := p.orm().DelId(&AuthOrg{}, id)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
