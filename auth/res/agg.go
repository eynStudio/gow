package res

import (
	"fmt"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/dddd/ddd"
)

type Res struct {
	Entity
	Id     GUID   `bson:"_id,omitempty"`
	Mc     string `Mc`
	Bz     string `Bz`
	Qz     int    `Qz`
	Icon   string `Icon`
	Uri    string `Uri`
	Opts   []Opt  `Opts`
	Params `Params`
}

func (p Res) ID() GUID       { return p.Id }
func (p Res) GetMc() string  { return p.Mc }
func (p Res) GetUri() string { return p.Uri }
func (p Res) GetQz() int     { return p.Qz }

func (p *Res) ReplaceOpt(opt Entity) {
	Slice(&p.Opts).ReplaceEntity(opt)
}
func (p *Res) DelOpt(id GUID) {
	Slice(&p.Opts).RemoveEntity(id)
}

type Opt struct {
	Id      GUID   `Id`
	Mc      string `Mc`
	Bz      string `Bz`
	Icon    string `Icon`
	Qz      int    `Qz`
	Visible bool   `Visible`
	Action  string `Action`
	Params  `Params`
}

func (p Opt) ID() GUID { return p.Id }

type ResAgg struct {
	AggBase
	root Res
}

func (p *ResAgg) RegistedCmds() []Cmd {
	return []Cmd{&SaveRes{}, &DelRes{}, &SaveResOpt{}, &DelResOpt{}}
}

func (p *ResAgg) HandleCmd(cmd Cmd) error {
	switch cmd := cmd.(type) {
	case *SaveRes:
		p.ApplyEvent((*ResSaved)(cmd))
	case *DelRes:
		p.ApplyEvent((*ResDeleted)(cmd))
	case *SaveResOpt:
		p.ApplyEvent((*ResOptSaved)(cmd))
	case *DelResOpt:
		p.ApplyEvent((*ResOptDeleted)(cmd))
	default:
		fmt.Println("ResAgg HandleCmd: no handler")
	}
	return nil
}

func (p *ResAgg) ApplyEvent(event Event) {
	switch evt := event.(type) {
	case *ResSaved:
		p.root = Res(*evt)
	case *ResDeleted:
		p.root = Res{}
	case *ResOptSaved:
		p.root.ReplaceOpt(evt.Opt)
	case *ResOptDeleted:
		p.root.DelOpt(evt.OptId)
	default:
		fmt.Println("ResAgg ApplyEvent: no handler")
	}

	p.StoreEvent(event)
}

func (p *ResAgg) Root() Entity {
	return &p.root
}
