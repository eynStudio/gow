package nodes

import (
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
	return h
}

func (p *RoleNode) Handler(c *gweb.Ctx) {
	handled := true
	switch c.Method {
	case "GET":
		//		p.Get(c)
	case "POST":
		p.Post(c)
	case "DELETE":
		c.OK()
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *RoleNode) Get(c *gweb.Ctx) {
	//	jbreak := c.Req.Header.Get("Authorization")
	//	if jbreak != "" {
	//		token := strings.Split(jbreak, " ")[1]
	//		if user, ok := p.LoginByToken(token); ok {
	//			c.Json(user)
	//			return
	//		}
	//	}
	//	c.Forbidden()
}

func (p *RoleNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetRoleTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.RoleRepo.Get(c.Get("id")))
	} else {
		c.Json(&role.Role{Id: p.RoleRepo.NewId()})
	}
	//	var login auth.Login
	//	c.Req.JsonBody(&login)
	//	log.Println(login)
	//	if user, ok := p.Login(&login); ok {
	//		c.Json(user)
	//		return
	//	}
	//	c.Json(auth.LoginErr{"登录失败"})
}
