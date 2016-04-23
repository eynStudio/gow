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

//func NewRegistedEventHandler(handler RegistedEventsHandler) RegistedEventsHandler {
//	di.Root.Apply(handler)
//	return handler
//}

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
	*ResRepo  `di`
	*OrgRepo  `di`
	*RoleRepo `di`
	*UserRepo `di`
}

func (p *AuthCtx) LoginByToken(token string) (*LoginOk, bool) {

	if token == "000000000000000000000000" {
		return &LoginOk{"000000000000000000000000", "超级管理员", p.GetResTree()}, true
	}

	//	if !bson.IsObjectIdHex(token) {
	//		return nil, false
	//	}

	//	u := p.UserRepo.Get(token).(*user.User)
	//	if u.Mc == "" {
	//		return nil, false
	//	} else {
	//		return &LoginOk{u.Id, u.Mc, p.GetResTreeByUser(u)}, true
	//	}

	return nil, false
}

func (p *AuthCtx) Login(login *Login) (*LoginOk, bool) {

	if login.UserName == "Super" {
		return &LoginOk{"000000000000000000000000", "超级管理员", p.GetResTree()}, true
	}

	//	authPass := false
	//	var u user.User

	//	u, authPass = p.UserRepo.GetUser(login.UserName, login.UserPwd)

	//	if authPass {
	//		return &LoginOk{u.Id, u.Mc, p.GetResTreeByUser(&u)}, true
	//	} else {
	//		return nil, false
	//	}
	return nil, false
}

func (p *AuthCtx) GetResTree() []*TreeNode {
	return BuildTree(p.ResRepo.All())
}

func (p *AuthCtx) GetOrgTree() []*TreeNode {
	return BuildTree(p.OrgRepo.All())
}

func (p *AuthCtx) GetRoleTree() []*TreeNode {
	return BuildTree(p.RoleRepo.All())
}
