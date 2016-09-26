package res

import (
	"log"

	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

func init() {
	log.Println("res.init")

	log.Println(di.Reg(&ResCtx{}))
}

type ResCtx struct {
	*orm.Orm `di:"*"`
}

func (p *ResCtx) All() (lst []AuthRes, err error) {
	err = p.Orm.AllJson(&lst)
	return
}

func (p *ResCtx) Save(m AuthRes) gobreak.IStatus {
	err := p.Orm.SaveJson(m.Id, m)
	if err != nil {
		log.Println(err)
	}
	return gobreak.NewStatusErr(err, "保存成功", "保存失败")
}
