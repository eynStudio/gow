package cms

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
)

type CateSaved Cate

func (p *CateSaved) ID() GUID { return p.Id }

type CateDeleted IdEvent

func (p *CateDeleted) ID() GUID { return p.Id }

type InfoSaved Info

func (p *InfoSaved) ID() GUID { return p.Id }

type InfoDeleted IdEvent

func (p *InfoDeleted) ID() GUID { return p.Id }
