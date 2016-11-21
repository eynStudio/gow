package auth

import (
	"github.com/eynstudio/gow/auth/group"
	"github.com/eynstudio/gox/utils"
)

type GroupRoles struct {
	Group group.AuthGroup
	Roles utils.TreeNodes
}
