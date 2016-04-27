package auth

import (
	. "github.com/eynstudio/gow/auth/org"
	"gopkg.in/mgo.v2/bson"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
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
