package nodes

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gow/auth/models"
	"github.com/eynstudio/gweb"
)

type ResNode struct {
	*gweb.Node
	*auth.AuthCtx `di`
}

func NewResNode() *ResNode {
	h := &ResNode{Node: gweb.NewNode("res", true)}
	h.NewParamNode("id", true)
	return h
}

func (p *ResNode) Handle(c *gweb.Ctx) {
	handled := true
	switch c.Method {
	case gweb.POST:
		p.Post(c)
	case gweb.PUT:
		p.Put(c)
	case gweb.DEL:
		p.Del(c)
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *ResNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetResTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.IResRepo.Get(c.Get("id")))
	} else {
		c.Json(&models.Res{Id: p.IResRepo.NewId()})
	}
}

func (p *ResNode) Put(c *gweb.Ctx) {
	var m models.Res
	c.Req.JsonBody(&m)
	p.SaveRes(m)
}

func (p *ResNode) Del(c *gweb.Ctx) {
	var id = GUID(c.Get("id"))
	p.IResRepo.Del(id)
}
