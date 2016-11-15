package group

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

func init() {
	log.Println(di.Reg(&GroupCtx{}))
}

type GroupCtx struct {
	*orm.Orm `di:"*"`
}

func (p *GroupCtx) Get(id gobreak.GUID) (m AuthGroup, ok bool) {
	ok = p.Orm.WhereId(id).GetJson2(&m)
	return
}

func (p *GroupCtx) All(orgid gobreak.GUID) (lst []AuthGroup, err error) {
	err = p.Orm.Where(`json->>'OrgId'=?`, orgid).AllJson(&lst).Err
	return
}

//func (p *GroupCtx) AllAsTree() (tree interface{}, err error) {
//	lst := make([]AuthOrg, 0)
//	if err = p.Orm.AllJson(&lst); err != nil {
//		return nil, err
//	}
//	return utils.BuildTree(lst), nil
//}

//func (p *GroupCtx) Save(m *AuthOrg) gobreak.IStatus {
//	err := p.Orm.SaveJson(m.Id, m)
//	if err != nil {
//		log.Println(err)
//	}
//	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
//}

//func (p *GroupCtx) Del(id gobreak.GUID) gobreak.IStatus {
//	err := p.Orm.DelId(&AuthOrg{}, id)
//	if err != nil {
//		log.Println(err)
//	}
//	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
//}
