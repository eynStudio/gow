package user

import (
	"fmt"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	. "github.com/eynstudio/gobreak/dddd/ddd"
)

type User struct {
	Entity
	Id     GUID       `bson:"_id,omitempty"`
	Mc     string     `Mc`
	Nc     string     `Nc` //昵称
	Img    string     `Img`
	Pwd    string     `Pwd`
	Xm     string     `Xm`
	Bz     string     `Bz`
	Lock   bool       `Lock`
	Auth   []UserAuth `Auth`
	Groups []GUID     `Groups`
}

func (p *User) AddGroup(gid GUID) {
	if -1 == Slice(&p.Groups).FindEntityIndex(gid) {
		p.Groups = append(p.Groups, gid)
	}
}
func (p *User) DelGroup(id GUID) {
	Slice(&p.Groups).RemoveEntity(id)
}

func NewUser() *User {
	return &User{Id: NewGuid(), Lock: false, Auth: make([]UserAuth, 0), Groups: make([]GUID, 0)}
}

type UserAuth struct {
	Mc string `Mc`
	Lx string `Lx`
}

func (p *User) AddAuth(mc, lx string) {
	p.Auth = append(p.Auth, UserAuth{mc, lx})
}

type UserAgg struct {
	AggBase
	root User
}

func (p *UserAgg) RegistedCmds() []Cmd {
	return []Cmd{&SaveUser{}, &DelUser{}, &SaveUserGroup{}, &DelUserGroup{}, &UpdateUserPwd{}}
}

func (p *UserAgg) HandleCmd(cmd Cmd) error {
	switch cmd := cmd.(type) {
	case *SaveUser:
		p.ApplyEvent((*UserSaved)(cmd))
	case *DelUser:
		p.ApplyEvent((*UserDeleted)(cmd))
	case *UpdateUserPwd:
		p.ApplyEvent((*UserPwdUpdated)(cmd))
	case *SaveUserGroup:
		p.ApplyEvent((*UserGroupSaved)(cmd))
	case *DelUserGroup:
		p.ApplyEvent((*UserGroupDeleted)(cmd))
	default:
		fmt.Println("UserAgg HandleCmd: no handler")
	}
	return nil
}

func (p *UserAgg) ApplyEvent(event Event) {
	switch evt := event.(type) {
	case *UserSaved:
		p.root = User(*evt)
	case *UserDeleted:
		p.root = User{}
	case *UserPwdUpdated:
		p.root.Pwd = evt.Pwd
	case *UserGroupSaved:
		p.root.AddGroup(evt.GroupId)
	case *UserGroupDeleted:
		p.root.DelGroup(evt.GroupId)
	}
	p.StoreEvent(event)
}

func (p *UserAgg) Root() Entity { return &p.root }
func (p *UserAgg) ID() GUID     { return p.root.ID() }
