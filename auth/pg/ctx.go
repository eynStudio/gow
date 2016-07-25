package pg

import (
	"log"

	"github.com/eynstudio/gobreak/di2"
	"github.com/eynstudio/gow/auth/user"
)

func init() {
	di2.Reg(&UserRepo{}, &ctx{})
}

type ctx struct {
	user.IUserRepo `di:"*"`
}

func (p *ctx) Hi() {
	log.Println("hi")
	p.Hi2()
}
