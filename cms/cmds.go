package cms

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
)

type SaveCate Cate

func (p *SaveCate) ID() GUID { return p.Id }

type DelCate IdCmd

func (p *DelCate) ID() GUID { return p.Id }

type SaveInfo Info

func (p *SaveInfo) ID() GUID { return p.Id }

type DelInfo IdCmd

func (p *DelInfo) ID() GUID { return p.Id }
