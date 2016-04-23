package auth

import (
	"fmt"

	"github.com/eynstudio/gow/auth/res"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	. "github.com/eynstudio/gobreak/ddd"
)

type ResRepo struct {
	MgoRepo
}

func NewResRepo() *ResRepo {
	return &ResRepo{NewMgoRepo("AuthRes", func() T { return &res.Res{} })}
}

type ResEventHandler struct {
	Repo *ResRepo `di`
}

func (p *ResEventHandler) RegistedEvents() []Event {
	return []Event{&res.ResSaved{}, &res.ResDeleted{}, &res.ResOptSaved{}, &res.ResOptDeleted{}}
}

func (p *ResEventHandler) HandleEvent(event Event) {
	switch event := event.(type) {
	case *res.ResSaved:
		p.Repo.Save(event.ID(), event)
	case *res.ResDeleted:
		p.Repo.Del(event.ID())
	case *res.ResOptSaved:
		m := p.Repo.Get(event.ID()).(*res.Res)
		m.ReplaceOpt(event.Opt)
		p.Repo.Save(m.Id, m)
	case *res.ResOptDeleted:
		m := p.Repo.Get(event.ID()).(*res.Res)
		m.DelOpt(event.OptId)
		p.Repo.Save(m.Id, m)
	default:
		fmt.Println("ResEventHandler: no handler")
	}
}
