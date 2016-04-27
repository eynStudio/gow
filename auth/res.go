package auth

import (
	"github.com/eynstudio/gow/auth/res"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
)

type ResRepo struct {
	MgoRepo
}

func NewResRepo() *ResRepo {
	return &ResRepo{NewMgoRepo("AuthRes", func() T { return &res.Res{} })}
}
