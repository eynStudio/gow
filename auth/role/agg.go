package role

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/dddd/ddd"
)

type Role struct {
	Entity
	Id  GUID         `bson:"_id,omitempty"`
	Mc  string       `Mc`
	Uri string       `Uri`
	Bz  string       `Bz`
	Qz  int          `Qz`
	Res []Permission `Res`
}

func (p Role) GetMc() string  { return p.Mc }
func (p Role) GetUri() string { return p.Uri }
func (p Role) GetQz() int     { return p.Qz }

type Permission struct {
	ResId GUID   `ResId`
	Opts  []GUID `Opts`
}

type RoleAgg struct {
	AggBase
	root Role
}

func (p *RoleAgg) RegistedCmds() []Cmd {
	return []Cmd{&SaveRole{}, &DelRole{}}
}

func (p *RoleAgg) HandleCmd(cmd Cmd) error {
	switch cmd := cmd.(type) {
	case *SaveRole:
		p.ApplyEvent((*RoleSaved)(cmd))
	case *DelRole:
		p.ApplyEvent((*RoleDeleted)(cmd))
	}
	return nil
}

func (p *RoleAgg) ApplyEvent(event Event) {
	switch evt := event.(type) {
	case *RoleSaved:
		p.root = Role(*evt)
	case *RoleDeleted:
		p.root = Role{}
	}
	p.StoreEvent(event)
}

func (p *RoleAgg) Root() Entity { return &p.root }
