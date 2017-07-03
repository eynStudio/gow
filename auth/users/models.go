package users

import (
	"time"

	. "github.com/eynstudio/gobreak"
)

var PtrAuthUser = &AuthUser{}

// UserZt 状态：0正常，1锁定
type UserStatus int

const (
	UserStatusDefault UserStatus = 0
	UserStatusLock    UserStatus = 1
)

type AuthUser struct {
	Id      GUID
	Name    string
	Nick    string
	Pwd     string
	Status  UserStatus
	Img     string
	Created time.Time
	Updated time.Time
	Auth    []UserAuth
	Groups  []GUID
}

// IsLock 是否锁定
func (au AuthUser) IsLock() bool { return au.Status == UserStatusLock }

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
	Name string
	Type string
}

func (p *AuthUser) AddAuth(name, _type string) {
	p.Auth = append(p.Auth, UserAuth{name, _type})
}

type UserLine struct {
	Id   GUID
	Name string
	Nick string
}
