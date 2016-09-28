package role

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
	"github.com/eynstudio/gox/utils"
)

func init() {
	log.Println(di.Reg(&RoleCtx{}))
}

type RoleCtx struct {
	*orm.Orm `di:"*"`
}

func (p *RoleCtx) Get(id gobreak.GUID) (m AuthRole, ok bool) {
	ok = p.Orm.WhereId(id).GetJson2(&m)
	return
}

func (p *RoleCtx) All() (lst []AuthRole, err error) {
	err = p.Orm.AllJson(&lst)
	return
}
func (p *RoleCtx) AllAsTree() (tree interface{}, err error) {
	var lst []AuthRole
	if err = p.Orm.AllJson(&lst); err != nil {
		return nil, err
	}
	return utils.BuildTree(lst), nil
}

func (p *RoleCtx) Save(m *AuthRole) gobreak.IStatus {
	err := p.Orm.SaveJson(m.Id, m)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
