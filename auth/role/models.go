package role

import (
	. "github.com/eynstudio/gobreak"
)

type AuthRole struct {
	Id        GUID
	Name      string
	Ns        string
	SortOrder int
	Res       []Permission
}

func NewRole() *AuthRole { return &AuthRole{Id: Guid()} }

func (p AuthRole) GetName() string   { return p.Name }
func (p AuthRole) GetNs() string     { return p.Ns }
func (p AuthRole) GetSortOrder() int { return p.SortOrder }
func (p AuthRole) GetId() GUID       { return p.Id }

type Permission struct {
	ResId GUID
	Opts  []GUID
}
