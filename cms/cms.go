package cms

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	. "github.com/eynstudio/gow/cms/models"
)

type CateRepo struct {
	MgoRepo
}
type InfoRepo struct {
	MgoRepo
}

func NewCateRepo() *CateRepo { return &CateRepo{NewMgoRepo("CmsCate", func() T { return &Cate{} })} }
func NewInfoRepo() *InfoRepo { return &InfoRepo{NewMgoRepo("CmsInfo", func() T { return &Info{} })} }
