package role

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
)

type SaveRole Role

func (p *SaveRole) ID() GUID { return p.Id }

type DelRole IdCmd

func (p *DelRole) ID() GUID { return p.Id }
