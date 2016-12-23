package auth

import (
	"log"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gow/auth/group"
	"github.com/eynstudio/gow/auth/res"
	"github.com/eynstudio/gow/auth/role"
	"github.com/eynstudio/gow/auth/users"
	"github.com/eynstudio/gow/auth/x/redissess"

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
	*users.UserCtx  `di:"*"`
}

func (ac AuthCtx) GetGroupRoles(id GUID) (m GroupRoles) {
	m.Group, _ = ac.GroupCtx.Get(id)
	m.Roles, _ = ac.RoleCtx.AllAsTree()
	return
}

func (ac AuthCtx) GetRoleRes(id GUID) (m RoleRes) {
	m.Role, _ = ac.RoleCtx.Get(id)
	m.Res, _ = ac.ResCtx.AllAsTree()
	return
}

func (ac AuthCtx) Login(req LoginReq) (resp LoginResp) {
	u, ok := ac.GetByMcPwd(req.Mc, SaltPwd(req.Pwd))

	if !ok || u.IsLock() {
		resp.ErrMsg("登录失败")
		return
	}

	resp.Token = Guid().String()
	//		resp.Id = u.Id
	redissess.SetSess(resp.Token, u.Id.String())
	resp.Ok()
	return
}
