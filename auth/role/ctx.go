package role

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/utils"
)

var Ctx = &RoleCtx{}

type RoleCtx struct {
}

func (p *RoleCtx) orm() *orm.Orm { return orm.GetOrmByName("auth") }

func (p *RoleCtx) Get(id gobreak.GUID) (m AuthRole, ok bool) {
	ok = p.orm().WhereId(id).GetJson2(&m)
	return
}

func (p *RoleCtx) All() (lst []AuthRole, err error) {
	err = p.orm().AllJson(&lst)
	return
}
func (p *RoleCtx) AllAsTree() (tree utils.TreeNodes, err error) {
	var lst []AuthRole
	if err = p.orm().AllJson(&lst); err != nil {
		return nil, err
	}
	return utils.BuildTree(lst), nil
}

func (p *RoleCtx) Save(m *AuthRole) gobreak.IStatus {
	err := p.orm().SaveJson(m.Id, m)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}

func (p *RoleCtx) Del(id gobreak.GUID) gobreak.IStatus {
	err := p.orm().DelId(&AuthRole{}, id)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
