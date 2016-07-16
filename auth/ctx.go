package auth

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/auth/domains"
	. "github.com/eynstudio/gow/auth/models"
	"github.com/eynstudio/gow/auth/repo"
)

var authctx IAuthCtx

func init() {

}

type IAuthCtx interface {
}

func Init() {
	repoRes := NewResRepo()
	di.MapAs(repoRes, (*repo.IResRepo)(nil)).Apply(repoRes.MgoRepo)

	repoRole := NewRoleRepo()
	di.MapAs(repoRole, (*repo.IRoleRepo)(nil)).Apply(repoRole.MgoRepo)

	repoOrg := NewOrgRepo()
	di.MapAs(repoOrg, (*repo.IOrgRepo)(nil)).Apply(repoOrg.MgoRepo)

	repoUser := NewUserRepo()
	di.MapAs(repoUser, (*repo.IUserRepo)(nil)).Apply(repoUser.MgoRepo)

	di.ApplyAndMap(&AuthCtx{})
}

type AuthCtx struct {
	repo.IResRepo  `di`
	repo.IOrgRepo  `di`
	repo.IRoleRepo `di`
	repo.IUserRepo `di`
	OnLogin        func(login *Login) (*LoginOk, bool)
	OnLoginByToken func(token string) (*LoginOk, bool)
}

func (p *AuthCtx) LoginByToken(token string) (*LoginOk, bool) { return p.OnLoginByToken(token) }
func (p *AuthCtx) Login(login *Login) (*LoginOk, bool)        { return p.OnLogin(login) }

func (p *AuthCtx) GetResTree() []*TreeNode {
	return BuildTree(p.IResRepo.All())
}

func (p *AuthCtx) GetOrgTree() []*TreeNode {
	return BuildTree(p.IOrgRepo.All())
}

func (p *AuthCtx) GetRoleTree() []*TreeNode {
	return BuildTree(p.IRoleRepo.All())
}

func (p *AuthCtx) GetUserPage() T {
	return p.IUserRepo.All()
}

func (p *AuthCtx) SaveOrg(m Org) {
	p.IOrgRepo.Save(m.Id, m)
}

func (p *AuthCtx) SaveRes(m Res) {
	p.IResRepo.Save(m.Id, m)
}
func (p *AuthCtx) SaveRole(m Role) {
	p.IRoleRepo.Save(m.Id, m)
}
func (p *AuthCtx) SaveUser(m User) {
	p.IUserRepo.Save(m.Id, m)
}

func (p *AuthCtx) GetRes(id GUID) (m *domains.ResAgg) {
	m = domains.NewResAgg(id)
	return
}

func (p *AuthCtx) GetRole(id GUID) (m *domains.RoleAgg) {
	m = domains.NewRoleAgg(id)
	return
}

func (p *AuthCtx) GetOrg(id GUID) (m *domains.OrgAgg) {
	m = domains.NewOrgAgg(id)
	return
}

func (p *AuthCtx) GetUser(id GUID) (m *domains.UserAgg) {
	m = domains.NewUserAgg(id)
	return
}

func (p *AuthCtx) GetNcImg(id GUID) (nc string, img string) {
	return p.GetUser(id).GetNcImg()
}

func (p *AuthCtx) CheckThenUpdatePwd(id GUID, pwd0, pwd1 string) Status {
	u := p.GetUser(id)
	u.CheckThenUpdatePwd(SaltPwd(pwd0), SaltPwd(pwd1))
	if u.NotErr() {

		return OkStatus
	}
	return u.GetStatus()
}

func (p *AuthCtx) UpdatePwd(id GUID, pwd string) Status {
	u := p.GetUser(id)
	u.UpdatePwd(SaltPwd(pwd))
	if u.NotErr() {
		return OkStatus
	}
	return u.GetStatus()
}
