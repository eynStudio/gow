package nodes

import (
	"fmt"

	"github.com/eynstudio/gobreak/dddd/cmdbus"
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gow/auth/role"
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

func (p *RoleNode) Handler(c *gweb.Ctx) {
	handled := true
	switch c.Method {
	case "GET":
		p.Get(c)
	case "POST":
		p.Post(c)
	case "PUT":
		p.Put(c)
	case "DELETE":
		c.OK()
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *RoleNode) Get(c *gweb.Ctx) {
}

func (p *RoleNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetRoleTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.RoleRepo.Get(c.Get("id")))
	} else {
		c.Json(&role.Role{Id: p.RoleRepo.NewId(), Res: []role.Permission{}})
	}
}
func (p *RoleNode) Put(c *gweb.Ctx) {
	var m role.SaveRole
	c.Req.JsonBody(&m)
	if err := cmdbus.Exec(&m); err != nil {
		fmt.Errorf("%#v", err)
	}
}
