package res

import (
	"fmt"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/dddd/ddd"
)

type Res struct {
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

func (p *ResAgg) ID() GUID     { return p.root.ID() }
func (p *ResAgg) Root() Entity { return &p.root }

func (p *ResAgg) RegistedCmds() []Cmd {
	return []Cmd{&SaveRes{}, &DelRes{}, &SaveResOpt{}, &DelResOpt{}}
}

func (p *ResAgg) HandleCmd(cmd Cmd) error {
	switch cmd := cmd.(type) {
	case *SaveRes:
		p.root = Res(*cmd)
		p.ApplyEvent((*ResSaved)(cmd))
	case *DelRes:
		p.root = Res{}
		p.ApplyEvent((*ResDeleted)(cmd))
	case *SaveResOpt:
		p.root.ReplaceOpt(cmd.Opt)
		p.ApplyEvent((*ResOptSaved)(cmd))
	case *DelResOpt:
		p.root.DelOpt(cmd.OptId)
		p.ApplyEvent((*ResOptDeleted)(cmd))
	default:
		fmt.Println("ResAgg HandleCmd: no handler")
	}
	return nil
}

func (p *ResAgg) ApplyEvent(event Event) {
	switch event.(type) {
	case *ResSaved:
	case *ResDeleted:
	case *ResOptSaved:
	case *ResOptDeleted:
	default:
		fmt.Println("ResAgg ApplyEvent: no handler")
	}

	p.StoreEvent(event)
}
