package res

import (
	"log"

	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

func init() {
	log.Println("res.init")

	log.Println(di.Reg(&ResCtx{Mc: "eyn"}))
}

type ResCtx struct {
	*orm.Orm `di:"*"`
	Mc       string
}

func (p *ResCtx) All() (lst []AuthRes, err error) {
	err = p.Orm.AllJson(&lst)
	return
}
