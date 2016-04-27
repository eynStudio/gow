package auth

import (
	"github.com/eynstudio/gow/auth/org"
	"github.com/eynstudio/gow/auth/res"
	"github.com/eynstudio/gow/auth/role"
	"github.com/eynstudio/gow/auth/user"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/dddd"
	"github.com/eynstudio/gobreak/di"
)

func Init() {
	repoRes := NewResRepo()
	di.Map(repoRes).Apply(repoRes.MgoRepo)
	dddd.Reg(&res.ResAgg{}, repoRes)

	repoRole := NewRoleRepo()
	di.Map(repoRole).Apply(repoRole.MgoRepo)
	dddd.Reg(&role.RoleAgg{}, repoRole)

	repoOrg := NewOrgRepo()
	di.Map(repoOrg).Apply(repoOrg.MgoRepo)
	dddd.Reg(&org.OrgAgg{}, repoOrg)

	repoUser := NewUserRepo()
	di.Map(repoUser).Apply(repoUser.MgoRepo)
	dddd.Reg(&user.UserAgg{}, repoUser)

	di.ApplyAndMap(&AuthCtx{})
}

type AuthCtx struct {
	*ResRepo       `di`
	*OrgRepo       `di`
	*RoleRepo      `di`
	*UserRepo      `di`
	OnLogin        func(login *Login) (*LoginOk, bool)
	OnLoginByToken func(token string) (*LoginOk, bool)
}

func (p *AuthCtx) LoginByToken(token string) (*LoginOk, bool) { return p.OnLoginByToken(token) }
func (p *AuthCtx) Login(login *Login) (*LoginOk, bool)        { return p.OnLogin(login) }

func (p *AuthCtx) GetResTree() []*TreeNode {
	return BuildTree(p.ResRepo.All())
}

func (p *AuthCtx) GetOrgTree() []*TreeNode {
	return BuildTree(p.OrgRepo.All())
}

func (p *AuthCtx) GetRoleTree() []*TreeNode {
	return BuildTree(p.RoleRepo.All())
}

func (p *AuthCtx) GetUserPage() T {
	return p.UserRepo.All()
}
