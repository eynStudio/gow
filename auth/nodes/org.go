package nodes

import (
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
	return h
}

func (p *OrgNode) Handler(c *gweb.Ctx) {
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

func (p *OrgNode) Get(c *gweb.Ctx) {
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

func (p *OrgNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetOrgTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.OrgRepo.Get(c.Get("id")))
	} else {
		c.Json(&org.Org{Id: p.OrgRepo.NewId()})
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
