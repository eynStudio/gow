package res

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/utils"
)

var Ctx = &ResCtx{}

type ResCtx struct {
}

func (p *ResCtx) orm() *orm.Orm { return orm.GetOrmByName("auth") }

func (p *ResCtx) Get(id gobreak.GUID) (m AuthRes, ok bool) {
	ok = p.orm().WhereId(id).GetJson2(&m)
	return
}

func (p *ResCtx) All() (lst []AuthRes, err error) {
	err = p.orm().AllJson(&lst)
	return
}
func (p *ResCtx) AllAsTree() (tree utils.TreeNodes, err error) {
	var lst []AuthRes
	if err = p.orm().AllJson(&lst); err != nil {
		return nil, err
	}
	return utils.BuildTree(lst), nil
}

func (p *ResCtx) Save(m *AuthRes) gobreak.IStatus {
	err := p.orm().SaveJson(m.Id, m)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}

func (p *ResCtx) Del(id gobreak.GUID) gobreak.IStatus {
	err := p.orm().DelId(&AuthRes{}, id)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
