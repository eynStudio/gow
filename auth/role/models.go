package role

import (
	. "github.com/eynstudio/gobreak"
)

type AuthRole struct {
	Id  GUID
	Mc  string
	Uri string
	Bz  string
	Qz  int
	Res []Permission
}

func (p AuthRole) GetMc() string  { return p.Mc }
func (p AuthRole) GetUri() string { return p.Uri }
func (p AuthRole) GetQz() int     { return p.Qz }

type Permission struct {
	ResId GUID
	Opts  []GUID
}
