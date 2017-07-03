package group

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
)

var Ctx = &GroupCtx{}

type GroupCtx struct {
}

func (p *GroupCtx) orm() *orm.Orm { return orm.GetOrmByName("auth") }

func (p *GroupCtx) Get(id gobreak.GUID) (m AuthGroup, ok bool) {
	ok = p.orm().WhereId(id).GetJson2(&m)
	return
}

func (p *GroupCtx) All(orgid gobreak.GUID) (lst []AuthGroup, err error) {
	err = p.orm().Where(`json->>'OrgId'=?`, orgid).AllJson(&lst).Err
	return
}

//func (p *GroupCtx) AllAsTree() (tree interface{}, err error) {
//	lst := make([]AuthOrg, 0)
//	if err = p.Orm.AllJson(&lst); err != nil {
//		return nil, err
//	}
//	return utils.BuildTree(lst), nil
//}

func (p *GroupCtx) Save(m *AuthGroup) gobreak.IStatus {
	err := p.orm().SaveJson(m.Id, m)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}

func (p *GroupCtx) Del(id gobreak.GUID) gobreak.IStatus {
	err := p.orm().DelId(&AuthGroup{}, id)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
