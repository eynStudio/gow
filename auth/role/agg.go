package role

//import (
//	. "github.com/eynstudio/gobreak"
//	. "github.com/eynstudio/gobreak/dddd/ddd"
//)

//type Role struct {
//	Id  GUID         `bson:"_id,omitempty"`
//	Mc  string       `Mc`
//	Uri string       `Uri`
//	Bz  string       `Bz`
//	Qz  int          `Qz`
//	Res []Permission `Res`
//}

//func (p Role) ID() GUID       { return p.Id }
//func (p Role) GetMc() string  { return p.Mc }
//func (p Role) GetUri() string { return p.Uri }
//func (p Role) GetQz() int     { return p.Qz }

//func NewRole() *Role { return &Role{Res: []Permission{}} }

//type Permission struct {
//	ResId GUID   `ResId`
//	Opts  []GUID `Opts`
//}

//type RoleAgg struct {
//	AggBase
//	root Role
//}

//func (p *RoleAgg) RegistedCmds() []Cmd {
//	return []Cmd{&SaveRole{}, &DelRole{}}
//}

//func (p *RoleAgg) HandleCmd(cmd Cmd) error {
//	switch cmd := cmd.(type) {
//	case *SaveRole:
//		p.root = Role(*cmd)
//		p.ApplyEvent((*RoleSaved)(cmd))
//	case *DelRole:
//		p.root = Role{}
//		p.ApplyEvent((*RoleDeleted)(cmd))
//	}
//	return nil
//}

//func (p *RoleAgg) ApplyEvent(event Event) {
//	switch event.(type) {
//	case *RoleSaved:
//	case *RoleDeleted:
//	}
//	p.StoreEvent(event)
//}

//func (p *RoleAgg) Root() Entity { return &p.root }
//func (p *RoleAgg) ID() GUID     { return p.root.ID() }
