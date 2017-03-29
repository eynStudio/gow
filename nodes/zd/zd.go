package zd

import (
	"github.com/eynstudio/gow/base/zd"

	"github.com/eynstudio/gox/di"
	"github.com/eynstudio/gweb"
)

type ZdNode struct {
	*gweb.Node
	*zd.ZdCtx `di:"*"`
}

func NewZdNode() *ZdNode {
	h := &ZdNode{Node: gweb.NewNode("zd/{res1}/{id1}/{res2}/{id2}", true)}
	di.Reg(h)
	return h
}

func (n *ZdNode) GetCate(c *gweb.Ctx) {
	//	all, _ := cn.CmsCtx.CateTree()
	c.Json(n.ZdCtx.All())
}
