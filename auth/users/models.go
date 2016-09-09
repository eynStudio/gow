package users

import (
	. "github.com/eynstudio/gobreak"
)

type AuthUser struct {
	Id     GUID
	Mc     string
	Pwd    string
	Bz     string
	Lock   bool
	Auth   []UserAuth
	Groups []GUID
}

func (p *AuthUser) AddGroup(gid GUID) {
	if -1 == Slice(&p.Groups).FindEntityIndex(gid) {
		p.Groups = append(p.Groups, gid)
	}
}
func (p *AuthUser) DelGroup(id GUID) {
	Slice(&p.Groups).RemoveEntity(id)
}

func NewUser() *AuthUser {
	return &AuthUser{Id: Guid(), Lock: false, Auth: make([]UserAuth, 0), Groups: make([]GUID, 0)}
}

type UserAuth struct {
	Mc string
	Lx string
}

func (p *AuthUser) AddAuth(mc, lx string) {
	p.Auth = append(p.Auth, UserAuth{mc, lx})
}
