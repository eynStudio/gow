package user

import (
	"log"

	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

var ctx *userCtx = &userCtx{}

func init() {
	di.Reg(ctx)
}

type userCtx struct {
	*orm.Orm `di:"*"`
}

func AddUser(mc, pwd string) {
	log.Println(mc, pwd)
	log.Println(ctx.Orm == nil)
}
