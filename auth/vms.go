package auth

import (
	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gow/auth/group"
	"github.com/eynstudio/gow/auth/role"
	"github.com/eynstudio/gox/utils"
)

type GroupRoles struct {
	Group group.AuthGroup
	Roles utils.TreeNodes
}

type RoleRes struct {
	Role role.AuthRole
	Res  utils.TreeNodes
}

type LoginReq struct {
	Mc  string
	Pwd string
	Lx  string
}

type LoginResp struct {
	Token string
	gobreak.Status
}

type GroupItem struct {
	group.AuthGroup
	Users int
}
