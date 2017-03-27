package org

import . "github.com/eynstudio/gobreak"

type AuthOrg struct {
	Id   GUID
	Mc   string
	Bz   string
	Ns   string
	Qz   int
	Args Params
}

func NewOrg() *AuthOrg          { return &AuthOrg{Id: Guid()} }
func (p AuthOrg) GetMc() string { return p.Mc }
func (p AuthOrg) GetNs() string { return p.Ns }
func (p AuthOrg) GetQz() int    { return p.Qz }
func (p AuthOrg) GetId() GUID   { return p.Id }
