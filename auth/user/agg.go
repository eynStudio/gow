package user

//import (
//	"fmt"
//	"time"

//	. "github.com/eynstudio/gobreak"
//	. "github.com/eynstudio/gobreak/db/mgo"
//	. "github.com/eynstudio/gobreak/dddd/ddd"
//)

//type User struct {
//	Id         GUID       `bson:"_id"`
//	Mc         string     `Mc`
//	Pwd        string     `Pwd`
//	Bz         string     `Bz`
//	Lock       bool       `Lock`
//	CreateTime time.Time  `CreateTime`
//	Auth       []UserAuth `Auth`
//	Groups     []GUID     `Groups`
//	Roles      []GUID     `Roles`
//}

//func (p *User) AddGroup(gid GUID) {
//	if -1 == Slice(&p.Groups).FindEntityIndex(gid) {
//		p.Groups = append(p.Groups, gid)
//	}
//}
//func (p *User) AddRole(rid GUID) {
//	if -1 == Slice(&p.Roles).FindEntityIndex(rid) {
//		p.Roles = append(p.Roles, rid)
//	}
//}
//func (p *User) DelGroup(id GUID) { Slice(&p.Groups).RemoveEntity(id) }
//func (p *User) DelRole(id GUID)  { Slice(&p.Roles).RemoveEntity(id) }
//func (p User) ID() GUID          { return p.Id }

//func NewUser() *User {
//	return &User{Id: NewGuid(), Lock: false, CreateTime: time.Now(),
//		Auth:   make([]UserAuth, 0),
//		Groups: make([]GUID, 0),
//	}
//}

//type UserAuth struct {
//	Mc string `Mc`
//	Lx string `Lx`
//}

//func (p *User) AddAuth(mc, lx string) {
//	p.Auth = append(p.Auth, UserAuth{mc, lx})
//}

//type UserAgg struct {
//	AggBase
//	root User
//}

//func (p *UserAgg) RegistedCmds() []Cmd {
//	return []Cmd{&SaveUser{}, &DelUser{}, &SaveUserGroup{}, &DelUserGroup{}, &UpdateUserPwd{}}
//}

//func (p *UserAgg) HandleCmd(cmd Cmd) error {
//	switch cmd := cmd.(type) {
//	case *SaveUser:
//		p.root = User(*cmd)
//		p.ApplyEvent((*UserSaved)(cmd))
//	case *DelUser:
//		p.root = User{}
//		p.ApplyEvent((*UserDeleted)(cmd))
//	case *UpdateUserPwd:
//		p.root.Pwd = cmd.Pwd
//		p.ApplyEvent((*UserPwdUpdated)(cmd))
//		//	case *UpdateUserNc:
//		//		p.root.Nc = cmd.Nc
//		//		p.ApplyEvent((*UserNcUpdated)(cmd))
//	case *SaveUserGroup:
//		p.root.AddGroup(cmd.GroupId)
//		p.ApplyEvent((*UserGroupSaved)(cmd))
//	case *DelUserGroup:
//		p.root.DelGroup(cmd.GroupId)
//		p.ApplyEvent((*UserGroupDeleted)(cmd))
//	default:
//		fmt.Println("UserAgg HandleCmd: no handler")
//	}
//	return nil
//}

//func (p *UserAgg) ApplyEvent(event Event) {
//	p.StoreEvent(event)
//}

//func (p *UserAgg) Root() Entity { return &p.root }
//func (p *UserAgg) ID() GUID     { return p.root.ID() }
