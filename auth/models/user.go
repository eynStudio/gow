package models

import (
	"time"

	. "github.com/eynstudio/gobreak"
)

type User struct {
	Id      GUID       `bson:"_id"`
	Mc      string     `Mc`  //名称，用户名
	Nc      string     `Nc`  //昵称
	Xm      string     `Xm`  //姓名
	Img     string     `Img` //头像
	Pwd     string     `Pwd`
	Bz      string     `Bz`
	Lock    bool       `Lock`
	Created time.Time  `Created`
	Updated time.Time  `Updated`
	Auths   []UserAuth `Auths`
	Groups  []GUID     `Groups`
	Roles   []GUID     `Roles`
	Ext     M          `Ext`
	Lx      int        `Lx`
}

//按Xm，Nc，Mc顺序查找第一个非空返回
func (p User) GetXmNcMc() string { return IfThenStr(p.Xm != "", p.Xm, p.GetNcMc()) }

//按Nc，Mc顺序查找第一个非空返回
func (p User) GetNcMc() string { return IfThenStr(p.Nc == "", p.Mc, p.Nc) }

func (p *User) AddGroup(gid GUID) {
	if -1 == Slice(&p.Groups).FindEntityIndex(gid) {
		p.Groups = append(p.Groups, gid)
	}
}
func (p *User) AddRole(rid GUID) {
	if -1 == Slice(&p.Roles).FindEntityIndex(rid) {
		p.Roles = append(p.Roles, rid)
	}
}
func (p *User) DelGroup(id GUID) { Slice(&p.Groups).RemoveEntity(id) }
func (p *User) DelRole(id GUID)  { Slice(&p.Roles).RemoveEntity(id) }
func (p User) ID() GUID          { return p.Id }

func NewUser(id GUID) *User {
	return &User{Id: id, Lock: false, Created: time.Now(),
		Auths:  make([]UserAuth, 0),
		Groups: make([]GUID, 0),
		Roles:  make([]GUID, 0),
	}
}

type UserAuth struct {
	Mc string `Mc`
	Lx string `Lx`
}

func (p *User) AddAuth(mc, lx string) {
	p.Auths = append(p.Auths, UserAuth{mc, lx})
}
