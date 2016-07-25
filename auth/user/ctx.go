package user

import (
	"log"

	"github.com/eynstudio/gobreak/di2"
	"github.com/eynstudio/gobreak/orm"
)

var ctx *userCtx = &userCtx{}

func init() {
	di2.Reg(ctx)
}

type userCtx struct {
	*orm.Orm `di:"*"`
}

func AddUser(mc, pwd string) {
	log.Println(mc, pwd)
	log.Println(ctx.Orm == nil)
}
