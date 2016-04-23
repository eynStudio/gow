package cms

import (
	"fmt"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	. "github.com/eynstudio/gobreak/ddd"
)

type CateRepo struct {
	MgoRepo
}
type InfoRepo struct {
	MgoRepo
}

func NewCateRepo() *CateRepo { return &CateRepo{NewMgoRepo("CmsCate", func() T { return &Cate{} })} }
func NewInfoRepo() *InfoRepo { return &InfoRepo{NewMgoRepo("CmsInfo", func() T { return &Info{} })} }

type CmsEventHandler struct {
	*CateRepo `di`
	*InfoRepo `di`
}

func (p *CmsEventHandler) RegistedEvents() []Event {
	return []Event{&CateSaved{}, &CateDeleted{},
		&InfoSaved{}, &InfoDeleted{},
	}
}

func (p *CmsEventHandler) HandleEvent(event Event) {
	switch event := event.(type) {
	case *CateSaved:
		p.CateRepo.Save(event.ID(), event)
	case *CateDeleted:
		p.CateRepo.Del(event.ID())
	case *InfoSaved:
		p.InfoRepo.Save(event.ID(), event)
	case *InfoDeleted:
		p.InfoRepo.Del(event.ID())
	default:
		fmt.Println("CmsEventHandler: no handler")
	}
}
