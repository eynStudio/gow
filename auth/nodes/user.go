package nodes

import (
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gow/auth/models"
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

func (p *UserNode) Handle(c *gweb.Ctx) {
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

func (p *UserNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetUserPage())
	} else if c.Scope.HasKey("id") {
		//		c.Json(p.UserRepo.Get(c.Get("id")))
	} else {
		c.Json(&models.User{Id: p.IUserRepo.NewId()})
	}
}

func (p *UserNode) Put(c *gweb.Ctx) {
	var m models.User
	c.Req.JsonBody(&m)
	p.SaveUser(m)
}
