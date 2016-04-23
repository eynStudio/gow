package res

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
)

type SaveRes Res

func (p *SaveRes) ID() GUID { return p.Id }

type DelRes IdCmd

func (p *DelRes) ID() GUID { return p.Id }

type SaveResOpt struct {
	Id  GUID
	Opt Opt
}

func (p *SaveResOpt) ID() GUID { return p.Id }

type DelResOpt struct {
	Id    GUID
	OptId GUID
}

func (p *DelResOpt) ID() GUID { return p.Id }
