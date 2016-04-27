package nodes

import (
	"fmt"

	"github.com/eynstudio/gobreak/dddd/cmdbus"
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gow/auth/org"
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

func (p *OrgNode) Get(c *gweb.Ctx) {

}

func (p *OrgNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetOrgTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.OrgRepo.Get(c.Get("id")))
	} else {
		c.Json(&org.Org{Id: p.OrgRepo.NewId()})
	}
}

func (p *OrgNode) Put(c *gweb.Ctx) {
	var m org.SaveOrg
	c.Req.JsonBody(&m)
	if err := cmdbus.Exec(&m); err != nil {
		fmt.Errorf("%#v", err)
	}

}
