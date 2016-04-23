package nodes

import (
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
	return h
}

func (p *UserNode) Handler(c *gweb.Ctx) {
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

func (p *UserNode) Get(c *gweb.Ctx) {
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

func (p *UserNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		//		c.Json(p.GetResTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.UserRepo.Get(c.Get("id")))
	} else {
		c.Json(&user.User{Id: p.UserRepo.NewId()})
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
