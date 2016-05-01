package nodes

import (
	"fmt"

	. "github.com/eynstudio/gobreak"
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

func (p *ResNode) Handle(c *gweb.Ctx) {
	handled := true
	switch c.Method {
	case gweb.POST:
		p.Post(c)
	case gweb.PUT:
		p.Put(c)
	case gweb.DEL:
		p.Del(c)
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *ResNode) Post(c *gweb.Ctx) {
	if c.JMethod() == "List" {
		c.Json(p.GetResTree())
	} else if c.Scope.HasKey("id") {
		c.Json(p.ResRepo.Get(c.Get("id")))
	} else {
		c.Json(&res.Res{Id: p.ResRepo.NewId()})
	}
}

func (p *ResNode) Put(c *gweb.Ctx) {
	var m res.SaveRes
	c.Req.JsonBody(&m)
	if err := cmdbus.Exec(&m); err != nil {
		fmt.Errorf("%#v", err)
	}
}

func (p *ResNode) Del(c *gweb.Ctx) {
	var m res.DelRes
	m.Id = GUID(c.Get("id"))
	if err := cmdbus.Exec(&m); err != nil {
		fmt.Errorf("%#v", err)
	}
}
