package cms

import (
	"fmt"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
)

type InfoAgg struct {
	*AggregateBase
	StateModel Info
}

func NewInfoAgg(id GUID) Aggregate {
	return &InfoAgg{
		AggregateBase: NewAggregateBase(id),
	}
}

func (p *InfoAgg) RegistedCmds() []Cmd {
	return []Cmd{&SaveInfo{}, &DelInfo{}}
}

func (p *InfoAgg) HandleCmd(cmd Cmd) error {
	switch cmd := cmd.(type) {
	case *SaveInfo:
		p.ApplyEvent((*InfoSaved)(cmd))
	case *DelInfo:
		p.ApplyEvent((*InfoDeleted)(cmd))
	default:
		fmt.Println("InfoAgg HandleCmd: no handler")
	}
	return nil
}

func (p *InfoAgg) ApplyEvent(event Event) {
	switch evt := event.(type) {
	case *InfoSaved:
		p.StateModel = Info(*evt)
	case *InfoDeleted:
		p.StateModel = Info{}
	}
	p.IncrementVersion()
	p.StoreEvent(event)
}

func (p *InfoAgg) GetSnapshot() T {
	return &p.StateModel
}
