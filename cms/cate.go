package cms

import (
	"fmt"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
)

type CateAgg struct {
	*AggregateBase
	StateModel Cate
}

func NewCateAgg(id GUID) Aggregate {
	return &CateAgg{
		AggregateBase: NewAggregateBase(id),
	}
}

func (p *CateAgg) RegistedCmds() []Cmd {
	return []Cmd{&SaveCate{}, &DelCate{}}
}

func (p *CateAgg) HandleCmd(cmd Cmd) error {
	switch cmd := cmd.(type) {
	case *SaveCate:
		p.ApplyEvent((*CateSaved)(cmd))
	case *DelCate:
		p.ApplyEvent((*CateDeleted)(cmd))
	default:
		fmt.Println("CateAgg HandleCmd: no handler")
	}
	return nil
}

func (p *CateAgg) ApplyEvent(event Event) {
	switch evt := event.(type) {
	case *CateSaved:
		p.StateModel = Cate(*evt)
	case *CateDeleted:
		p.StateModel = Cate{}
	}
	p.IncrementVersion()
	p.StoreEvent(event)
}

func (p *CateAgg) GetSnapshot() T {
	return &p.StateModel
}
