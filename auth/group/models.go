package group

import . "github.com/eynstudio/gobreak"

type AuthGroup struct {
	Id    GUID
	OrgId GUID
	Mc    string
	Bz    string
	Roles []GUID
	Args  Params
}

func (p AuthGroup) ID() GUID { return p.Id }

func NewGroup(orgid GUID) *AuthGroup { return &AuthGroup{Id: Guid(), OrgId: orgid, Args: []KeyValue{}} }
