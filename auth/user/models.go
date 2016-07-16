package user

import (
	"time"

	. "github.com/eynstudio/gobreak"
)

type AuthUser struct {
	Id      GUID       `Id`
	Mc      string     `Mc` //名称，用户名
	Pwd     string     `Pwd`
	Xm      string     `Xm`    //姓名
	Nc      string     `Nc`    //昵称
	Img     string     `Img`   //头像
	Phone   string     `Phone` //手机
	Bz      string     `Bz`
	Lock    bool       `Lock`
	Created time.Time  `Created`
	Updated time.Time  `Updated`
	Auths   []UserAuth `Auths`
	Groups  []GUID     `Groups`
	Roles   []GUID     `Roles`
	Ext     M          `Ext`
}

func NewUser(id GUID) *AuthUser {
	return &AuthUser{Id: Guid(), Lock: false, Created: time.Now(),
		Auths:  make([]UserAuth, 0),
		Groups: make([]GUID, 0),
		Roles:  make([]GUID, 0),
	}
}

type UserAuth struct {
	Mc string `Mc`
	Lx string `Lx`
}

func (p *AuthUser) AddAuth(mc, lx string) {
	p.Auths = append(p.Auths, UserAuth{mc, lx})
}
