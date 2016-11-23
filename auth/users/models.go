package users

import (
	"time"

	. "github.com/eynstudio/gobreak"
)

type AuthUser struct {
	Id      GUID
	Mc      string
	Nc      string
	Pwd     string
	Bz      string
	Zt      int //状态：0正常，1锁定，2存档，-1删除
	Img     string
	Created time.Time
	Updated time.Time
	Auth    []UserAuth
	Groups  []GUID
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
	return &AuthUser{Id: Guid(), Auth: make([]UserAuth, 0), Created: time.Now(), Updated: time.Now(), Groups: make([]GUID, 0)}
}

type UserAuth struct {
	Mc string
	Lx string
}

func (p *AuthUser) AddAuth(mc, lx string) {
	p.Auth = append(p.Auth, UserAuth{mc, lx})
}

type UserLine struct {
	Id GUID
	Mc string
	Nc string
	Lx int
}
