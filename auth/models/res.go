package models

import (
	. "github.com/eynstudio/gobreak"
)

type Res struct {
	Id     GUID   `bson:"_id"`
	Mc     string `Mc`
	Bz     string `Bz`
	Qz     int    `Qz`
	Icon   string `Icon`
	Uri    string `Uri`
	Opts   []Opt  `Opts`
	Params `Params`
}

func (p Res) ID() GUID       { return p.Id }
func (p Res) GetMc() string  { return p.Mc }
func (p Res) GetUri() string { return p.Uri }
func (p Res) GetQz() int     { return p.Qz }

func (p *Res) ReplaceOpt(opt Entity) {
	Slice(&p.Opts).ReplaceEntity(opt)
}
func (p *Res) DelOpt(id GUID) {
	Slice(&p.Opts).RemoveEntity(id)
}

type Opt struct {
	Id      GUID   `Id`
	Mc      string `Mc`
	Bz      string `Bz`
	Icon    string `Icon`
	Qz      int    `Qz`
	Visible bool   `Visible`
	Action  string `Action`
	Params  `Params`
}

func (p Opt) ID() GUID { return p.Id }
