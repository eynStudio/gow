package nodes

import (
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gow/auth/models"
	"github.com/eynstudio/gweb"
)

type OrgNode struct {
	*gweb.Node
	*auth.AuthCtx `di`
}

func NewOrgNode() *OrgNode {
	h := &OrgNode{Node: gweb.NewNode("org", true)}
	h.NewParamNode("id", true)
	return h
}

func (p *OrgNode) Handle(c *gweb.Ctx) {
	handled := true
	switch c.Method {
	case gweb.POST:
		p.Post(c)
	case gweb.PUT:
		p.Put(c)
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *OrgNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetOrgTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.IOrgRepo.Get(c.Get("id")))
	} else {
		c.Json(&models.Org{Id: p.IOrgRepo.NewId()})
	}
}

func (p *OrgNode) Put(c *gweb.Ctx) {
	var m models.Org
	c.Req.JsonBody(&m)
	p.SaveOrg(m)
}
