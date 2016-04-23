package auth

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	. "github.com/eynstudio/gobreak/ddd"

	"github.com/eynstudio/gow/auth/role"
)

type RoleRepo struct {
	MgoRepo
}

func NewRoleRepo() *RoleRepo {
	return &RoleRepo{NewMgoRepo("AuthRole", func() T { return &role.Role{} })}
}

type RoleEventHandler struct {
	Repo *RoleRepo `di`
}

func (p *RoleEventHandler) RegistedEvents() []Event {
	return []Event{&role.RoleSaved{}, &role.RoleDeleted{}}
}

func (p *RoleEventHandler) HandleEvent(event Event) {
	switch event := event.(type) {
	case *role.RoleSaved:
		p.Repo.Save(event.ID(), event)
	case *role.RoleDeleted:
		p.Repo.Del(event.ID())
	}
}
