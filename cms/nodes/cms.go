package cms

import (
	"log"

	"github.com/eynstudio/gow/cms/cms"

	"github.com/eynstudio/gox/di"
	"github.com/eynstudio/gweb"
)

type CmsNode struct {
	*gweb.Node
}

func NewCmsNode() *CmsNode {
	h := &CmsNode{Node: gweb.NewNode("cms/{res1}/{id1}/{res2}/{id2}", true)}
	di.Reg(h)
	return h
}

func (cn *CmsNode) GetCate(c *gweb.Ctx) {
	all, _ := cms.Ctx.CateTree()
	c.Json(all.Nodes)
}

func (cn *CmsNode) DeleteCate(c *gweb.Ctx) {
	id1 := c.Scope.GetGuid("id1")
	log.Println(id1)
	cms.Ctx.Orm().WhereId(id1).Del(&cms.CmsInfo{})
	c.Json(nil)
}

func (cn *CmsNode) GetCateId1(c *gweb.Ctx) {
	id1 := c.Scope.GetGuid("id1")
	c.Json(cms.Ctx.GetCate(id1))
}

func (cn *CmsNode) PutCate(c *gweb.Ctx, m *cms.CmsInfo) {
	m.Uid = c.Uid()
	c.Json(cms.Ctx.SaveCate(m))
}

func (cn *CmsNode) GetCateInfo(c *gweb.Ctx) {
	id1 := c.Scope.GetGuid("id1")
	c.Json(cms.Ctx.GetCateInfo(id1))
}

func (cn *CmsNode) GetInfoId1(c *gweb.Ctx) {
	id1 := c.Scope.GetGuid("id1")
	c.Json(cms.Ctx.GetInfo(id1))
}
func (cn *CmsNode) PutInfo(c *gweb.Ctx, m *cms.CmsInfo) {
	m.Uid = c.Uid()
	c.Json(cms.Ctx.SaveInfo(m))
}

func (cn *CmsNode) PostFile(c *gweb.Ctx) {
	c.Json(cms.UploadFile(c.Request))
}
func (cn *CmsNode) PostImg(c *gweb.Ctx) {
	c.Text(cms.UploadImg(c.Request).Url)
}
