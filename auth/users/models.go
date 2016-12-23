package users

import (
	"time"

	. "github.com/eynstudio/gobreak"
)

var PtrAuthUser = &AuthUser{}

// UserZt 状态：0正常，1锁定
type UserZt int

const (
	UserZtZc UserZt = 0
	UserZtSd UserZt = 1
)

type AuthUser struct {
	Id      GUID
	Mc      string
	Nc      string
	Pwd     string
	Bz      string
	Zt      UserZt
	Img     string
	Created time.Time
	Updated time.Time
	Auth    []UserAuth
	Groups  []GUID
}

// IsLock 是否锁定
func (au AuthUser) IsLock() bool { return au.Zt == UserZtSd }

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
