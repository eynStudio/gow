package auth

import (
	"fmt"

	. "github.com/eynstudio/gow/auth/org"
	"gopkg.in/mgo.v2/bson"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	. "github.com/eynstudio/gobreak/ddd"
)

type OrgRepo struct {
	MgoRepo
}

func NewOrgRepo() *OrgRepo { return &OrgRepo{NewMgoRepo("AuthOrg", func() T { return &Org{} })} }

func (p *OrgRepo) GetByUri(uri string) (m Org) {
	p.C(nil).Find(bson.M{"Uri": uri}).One(&m)
	return
}

func (p *OrgRepo) GetGroupByMc(uri, groupMc string) (m *Group) {
	var org Org
	p.C(nil).Find(bson.M{"Uri": uri}).One(&org)
	for _, g := range org.Groups {
		if g.Mc == groupMc {
			return &g
		}
	}
	return nil
}

type OrgEventHandler struct {
	Repo *OrgRepo `di`
}

func (p *OrgEventHandler) RegistedEvents() []Event {
	return []Event{&OrgSaved{}, &OrgDeleted{}, &OrgGroupSaved{}, &OrgGroupDeleted{}}
}

func (p *OrgEventHandler) HandleEvent(event Event) {
	switch event := event.(type) {
	case *OrgSaved:
		p.Repo.Save(event.ID(), event)
	case *OrgDeleted:
		p.Repo.Del(event.ID())
	case *OrgGroupSaved:
		m := p.Repo.Get(event.ID()).(*Org)
		m.ReplaceGroup(event.Group)
		p.Repo.Save(m.Id, m)
	case *OrgGroupDeleted:
		m := p.Repo.Get(event.ID()).(*Org)
		m.DelGroup(event.GroupId)
		p.Repo.Save(m.Id, m)
	default:
		fmt.Println("OrgEventHandler: no handler")
	}
}
