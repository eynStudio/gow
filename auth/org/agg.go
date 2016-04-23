package org

import (
	"fmt"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/dddd/ddd"
)

type Org struct {
	Entity
	Id     GUID    `bson:"_id,omitempty"`
	Mc     string  `Mc`
	Dm     string  `Dm`
	Uri    string  `Uri`
	Qz     int     `bson:"Qz" json:"Qz,string"`
	Groups []Group `Groups`
	Params `Params`
}

func (p Org) GetMc() string  { return p.Mc }
func (p Org) GetUri() string { return p.Uri }
func (p Org) GetQz() int     { return p.Qz }

func (p *Org) ReplaceGroup(group Entity) {
	Slice(&p.Groups).ReplaceEntity(group)
}
func (p *Org) GetGroup(groupId GUID) Group {
	return Slice(&p.Groups).FindEntity(groupId).(Group)
}
func (p *Org) DelGroup(id GUID) {
	Slice(&p.Groups).RemoveEntity(id)
}

type Group struct {
	Id     GUID   `Id`
	Mc     string `Mc`
	Bz     string `Bz`
	Roles  []GUID `Roles`
	Params `Params`
}

func (p Group) ID() GUID { return p.Id }

func NewGroup(id GUID) *Group {
	return &Group{Id: id, Params: []KeyValue{}}
}

type OrgAgg struct {
	AggBase
	root Org
}

func (p *OrgAgg) RegistedCmds() []Cmd {
	return []Cmd{&SaveOrg{}, &DelOrg{}, &SaveOrgGroup{}, &DelOrgGroup{}}
}

func (p *OrgAgg) HandleCmd(cmd Cmd) error {
	switch cmd := cmd.(type) {
	case *SaveOrg:
		p.ApplyEvent((*OrgSaved)(cmd))
	case *DelOrg:
		p.ApplyEvent((*OrgDeleted)(cmd))
	case *SaveOrgGroup:
		p.ApplyEvent((*OrgGroupSaved)(cmd))
	case *DelOrgGroup:
		p.ApplyEvent((*OrgGroupDeleted)(cmd))
	default:
		fmt.Println("OrgAgg HandleCmd: no handler")
	}
	return nil
}

func (p *OrgAgg) ApplyEvent(event Event) {
	switch evt := event.(type) {
	case *OrgSaved:
		p.root = Org(*evt)
	case *OrgDeleted:
		p.root = Org{}
	case *OrgGroupSaved:
		p.root.ReplaceGroup(evt.Group)
	case *OrgGroupDeleted:
		p.root.DelGroup(evt.GroupId)
	}
	p.StoreEvent(event)
}

func (p *OrgAgg) Root() Entity { return &p.root }
