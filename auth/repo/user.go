package repo

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	"github.com/eynstudio/gow/auth/models"
)

type IUserRepo interface {
	MgoRepo
	GetById(id GUID) (m *models.User, err error)
	UpdateNc(uid GUID, nc string)
	GetUserByMcPwd(mc, pwd string) (u *models.User, ok bool)
	GetUserByMc(mc string) (u *models.User, ok bool)
}
