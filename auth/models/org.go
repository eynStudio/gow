package models

import (
	. "github.com/eynstudio/gobreak"
)

type Org struct {
	Id     GUID    `bson:"_id"`
	Mc     string  `Mc`
	Dm     string  `Dm`
	Uri    string  `Uri`
	Qz     int     `Qz`
	Groups []Group `Groups`
	Params `Params`
}

func (p Org) ID() GUID       { return p.Id }
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
