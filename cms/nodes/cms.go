package cms

import (
	"github.com/eynstudio/gow/cms/cms"

	"github.com/eynstudio/gox/di"
	"github.com/eynstudio/gweb"
)

type CmsNode struct {
	*gweb.Node
	*cms.CmsCtx `di:"*"`
}

func NewCmsNode() *CmsNode {
	h := &CmsNode{Node: gweb.NewNode("cms/{res1}/{id1}/{res2}/{id2}", true)}
	di.Reg(h)
	return h
}

func (cn *CmsNode) GetCate(c *gweb.Ctx) {
	all, _ := cn.CmsCtx.CateTree()
	c.Json(all.Nodes)
}

func (cn *CmsNode) GetCateId1(c *gweb.Ctx) {
	id1 := c.Scope.GetGuid("id1")
	c.Json(cn.CmsCtx.GetCate(id1))
}

func (cn *CmsNode) PutCate(c *gweb.Ctx, m *cms.CmsInfo) {
	m.Uid = c.Uid()
	c.Json(cn.CmsCtx.SaveCate(m))
}
