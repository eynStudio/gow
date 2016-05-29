package models

import (
	. "github.com/eynstudio/gobreak"
)

type Role struct {
	Id  GUID         `bson:"_id"`
	Mc  string       `Mc`
	Uri string       `Uri`
	Bz  string       `Bz`
	Qz  int          `Qz`
	Res []Permission `Res`
}

func (p Role) ID() GUID       { return p.Id }
func (p Role) GetMc() string  { return p.Mc }
func (p Role) GetUri() string { return p.Uri }
func (p Role) GetQz() int     { return p.Qz }

func NewRole() *Role { return &Role{Res: []Permission{}} }

type Permission struct {
	ResId GUID   `ResId`
	Opts  []GUID `Opts`
}
