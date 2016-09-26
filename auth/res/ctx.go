package res

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
	"github.com/eynstudio/gox/utils"
)

func init() {
	log.Println("res.init")

	log.Println(di.Reg(&ResCtx{}))
}

type ResCtx struct {
	*orm.Orm `di:"*"`
}

func (p *ResCtx) Get(id gobreak.GUID) (m AuthRes, ok bool) {
	ok = p.Orm.WhereId(id).GetJson2(&m)
	return
}

func (p *ResCtx) All() (lst []AuthRes, err error) {
	err = p.Orm.AllJson(&lst)
	return
}
func (p *ResCtx) AllAsTree() (tree interface{}, err error) {
	var lst []AuthRes
	if err = p.Orm.AllJson(&lst); err != nil {
		return nil, err
	}
	return utils.BuildTree(lst), nil
}

func (p *ResCtx) Save(m *AuthRes) gobreak.IStatus {
	err := p.Orm.SaveJson(m.Id, m)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
