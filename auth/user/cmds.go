package user

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
)

type SaveUser User

func (p *SaveUser) ID() GUID { return p.Id }

type DelUser IdCmd

func (p *DelUser) ID() GUID { return p.Id }

type SaveUserGroup struct {
	Id      GUID
	GroupId GUID
}

func (p *SaveUserGroup) ID() GUID { return p.Id }

type DelUserGroup SaveUserGroup

func (p *DelUserGroup) ID() GUID { return p.Id }

type UpdateUserPwd struct {
	Id  GUID
	Pwd string
}

func (p *UpdateUserPwd) ID() GUID { return p.Id }
