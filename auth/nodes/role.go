package nodes

import (
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gow/auth/models"
	"github.com/eynstudio/gweb"
)

type RoleNode struct {
	*gweb.Node
	*auth.AuthCtx `di`
}

func NewRoleNode() *RoleNode {
	h := &RoleNode{Node: gweb.NewNode("role", true)}
	h.NewParamNode("id", true)
	return h
}

func (p *RoleNode) Handle(c *gweb.Ctx) {
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

func (p *RoleNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetRoleTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.IRoleRepo.Get(c.Get("id")))
	} else {
		c.Json(&models.Role{Id: p.IRoleRepo.NewId(), Res: []models.Permission{}})
	}
}
func (p *RoleNode) Put(c *gweb.Ctx) {
	var m models.Role
	c.Req.JsonBody(&m)
	p.SaveRole(m)
}
