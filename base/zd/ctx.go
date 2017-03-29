package zd

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

func init() {
	Must(di.Reg(&ZdCtx{}))
}

type ZdCtx struct {
	*orm.Orm `di:"*"`
}

func (ctx *ZdCtx) All() (lst []XtZd) {
	ctx.Orm.AllJson(&lst)
	return
}

//func (c *ZdCtx) CateTree() (tree *CateTree, err error) {
//	var lst []XtZd
//	s := c.Orm.Order("json->'Ns' desc", "json->'Qz'").AllJson(&lst)
//	if s.IsErr() {
//		return nil, s.Err
//	}
//	for _, it := range lst {
//		log.Println(it.Ns, it.Qz, it.Mc, it.GetUri())
//	}
//	tree = NewCateTree()
//	log.Println(lst)
//	tree.Build(lst)
//	return tree, nil
//}
