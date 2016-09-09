package org

import . "github.com/eynstudio/gobreak"

type AuthOrg struct {
	Id     GUID
	Mc     string
	Dm     string
	Uri    string
	Qz     int
	Groups []Group
	Params
}

func (p AuthOrg) GetMc() string  { return p.Mc }
func (p AuthOrg) GetUri() string { return p.Uri }
func (p AuthOrg) GetQz() int     { return p.Qz }

func (p *AuthOrg) ReplaceGroup(group Entity)   { Slice(&p.Groups).ReplaceEntity(group) }
func (p *AuthOrg) GetGroup(groupId GUID) Group { return Slice(&p.Groups).FindEntity(groupId).(Group) }
func (p *AuthOrg) DelGroup(id GUID)            { Slice(&p.Groups).RemoveEntity(id) }

type Group struct {
	Id     GUID
	Mc     string
	Bz     string
	Roles  []GUID
	Params []KeyValue
}

func (p Group) ID() GUID { return p.Id }

func NewGroup(id GUID) *Group { return &Group{Id: id, Params: []KeyValue{}} }
