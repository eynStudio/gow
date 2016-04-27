package nodes

import (
	"log"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gweb"
)

type AuthNode struct {
	*gweb.Node
	*auth.AuthCtx `di`
}

func NewAuthNode() *AuthNode {
	h := &AuthNode{Node: gweb.NewNode("auth", false)}
	h.AddNode(NewResNode())
	h.AddNode(NewRoleNode())
	h.AddNode(NewUserNode())
	h.AddNode(NewOrgNode())
	return h
}

func (p *AuthNode) Handle(c *gweb.Ctx) {
	handled := true
	switch c.Method {
	case "GET":
		p.Get(c)
	case "POST":
		p.Post(c)
	case "DELETE":
		c.OK()
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *AuthNode) Get(c *gweb.Ctx) {
	if c.HasToken() {
		if user, ok := p.LoginByToken(c.Token); ok {
			c.Json(user)
			return
		}
	}
	//	jbreak := c.Req.Header.Get("Authorization")
	//	if jbreak != "" {
	//		token := strings.Split(jbreak, " ")[1]
	//		if user, ok := p.LoginByToken(token); ok {
	//			c.Json(user)
	//			return
	//		}
	//	}
	c.Forbidden()
}

func (p *AuthNode) Post(c *gweb.Ctx) {
	var login auth.Login
	c.Req.JsonBody(&login)
	log.Println(login)
	if user, ok := p.Login(&login); ok {
		c.Json(user)
		return
	}
	c.Json(auth.LoginErr{"登录失败"})
}
