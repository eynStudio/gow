package user

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
)

type UserSaved User

func (p *UserSaved) ID() GUID { return p.Id }

type UserDeleted IdEvent

func (p *UserDeleted) ID() GUID { return p.Id }

type UserGroupSaved struct {
	Id      GUID
	GroupId GUID
}

func (p *UserGroupSaved) ID() GUID { return p.Id }

type UserGroupDeleted UserGroupSaved

func (p *UserGroupDeleted) ID() GUID { return p.Id }

type UserPwdUpdated struct {
	Id  GUID
	Pwd string
}

func (p *UserPwdUpdated) ID() GUID { return p.Id }

type UserNcUpdated struct {
	Id GUID
	Nc string
}

func (p *UserNcUpdated) ID() GUID { return p.Id }
