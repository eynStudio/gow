package auth

import (
	"log"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gow/auth/group"
	"github.com/eynstudio/gow/auth/role"

	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

func init() {
	log.Println(di.Reg(&AuthCtx{}))
}

type AuthCtx struct {
	*orm.Orm        `di:"*"`
	*group.GroupCtx `di:"*"`
	*role.RoleCtx   `di:"*"`
}

func (p AuthCtx) GetGroupRoles(id GUID) (m GroupRoles) {
	m.Group, _ = p.GroupCtx.Get(id)
	m.Roles, _ = p.RoleCtx.AllAsTree()
	return
}
