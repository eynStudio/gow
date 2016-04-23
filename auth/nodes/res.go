package nodes

import (
	"fmt"

	"github.com/eynstudio/gobreak/dddd/cmdbus"
	"github.com/eynstudio/gow/auth"
	"github.com/eynstudio/gow/auth/res"
	"github.com/eynstudio/gweb"
)

type ResNode struct {
	*gweb.Node
	*auth.AuthCtx `di`
}

func NewResNode() *ResNode {
	h := &ResNode{Node: gweb.NewNode("res", true)}
	h.NewParamNode("id", true)
	return h
}

func (p *ResNode) Handler(c *gweb.Ctx) {
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

func (p *ResNode) Get(c *gweb.Ctx) {
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

func (p *ResNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetResTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.ResRepo.Get(c.Get("id")))
	} else {
		c.Json(&res.Res{Id: p.ResRepo.NewId()})
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

func (p *ResNode) Put(c *gweb.Ctx) {
	var m res.SaveRes
	c.Req.JsonBody(&m)
	if err := cmdbus.Exec((*res.SaveRes)(&m)); err != nil {
		fmt.Errorf("%#v", err)
	}

}
