package auth

import (
	"log"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gow/auth/group"
	"github.com/eynstudio/gow/auth/res"
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
	*res.ResCtx     `di:"*"`
}

func (p AuthCtx) GetGroupRoles(id GUID) (m GroupRoles) {
	m.Group, _ = p.GroupCtx.Get(id)
	m.Roles, _ = p.RoleCtx.AllAsTree()
	return
}

func (p AuthCtx) GetRoleRes(id GUID) (m RoleRes) {
	m.Role, _ = p.RoleCtx.Get(id)
	m.Res, _ = p.ResCtx.AllAsTree()
	return
}
