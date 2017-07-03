package org

import . "github.com/eynstudio/gobreak"

type AuthOrg struct {
	Id        GUID
	Name      string
	Ns        string
	SortOrder int
	Args      Params
}

func NewOrg() *AuthOrg              { return &AuthOrg{Id: Guid()} }
func (p AuthOrg) GetName() string   { return p.Name }
func (p AuthOrg) GetNs() string     { return p.Ns }
func (p AuthOrg) GetSortOrder() int { return p.SortOrder }
func (p AuthOrg) GetId() GUID       { return p.Id }
