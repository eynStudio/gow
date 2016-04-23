package nodes

import (
	"fmt"

	"github.com/eynstudio/gobreak/dddd/cmdbus"
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gow/auth/user"
	"github.com/eynstudio/gweb"
)

type UserNode struct {
	*gweb.Node
	*auth.AuthCtx `di`
}

func NewUserNode() *UserNode {
	h := &UserNode{Node: gweb.NewNode("user", true)}
	h.NewParamNode("id", true)
	return h
}

func (p *UserNode) Handler(c *gweb.Ctx) {
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

func (p *UserNode) Get(c *gweb.Ctx) {

}

func (p *UserNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		//		c.Json(p.GetResTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.UserRepo.Get(c.Get("id")))
	} else {
		c.Json(&user.User{Id: p.UserRepo.NewId()})
	}
}

func (p *UserNode) Put(c *gweb.Ctx) {
	var m user.SaveUser
	c.Req.JsonBody(&m)
	if err := cmdbus.Exec(&m); err != nil {
		fmt.Errorf("%#v", err)
	}
}
