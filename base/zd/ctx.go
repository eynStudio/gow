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

func (c *ZdCtx) All() (lst []XtZd) {
	c.Orm.AllJson(&lst)
	return
}

func (c *ZdCtx) GetZd(id GUID) (m XtZd) {
	if id.IsEmpty() {
		return NewXtZd()
	}
	c.Orm.WhereId(id).GetJson2(&m)
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
