package role

import (
	. "github.com/eynstudio/gobreak"
)

type AuthRole struct {
	Id  GUID
	Mc  string
	Ns  string
	Bz  string
	Qz  int
	Res []Permission
}

func NewRole() *AuthRole { return &AuthRole{Id: Guid()} }

func (p AuthRole) GetMc() string { return p.Mc }
func (p AuthRole) GetNs() string { return p.Ns }
func (p AuthRole) GetQz() int    { return p.Qz }
func (p AuthRole) GetId() GUID   { return p.Id }

type Permission struct {
	ResId GUID
	Opts  []GUID
}
