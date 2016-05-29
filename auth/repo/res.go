package repo

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	"github.com/eynstudio/gow/auth/models"
)

type IResRepo interface {
	MgoRepo
	GetById(id GUID) (m *models.Res, err error)
}
