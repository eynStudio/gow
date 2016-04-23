package cms

import (
	"time"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/db/mgo"
)

type Xx struct {
	Mc          string `Mc`
	Mc2         string `Mc2`
	Qz          int    `Qz`
	Lx          string `Lx`
	Nr          string `Nr`
	ImgUrl      string `ImgUrl`
	TargetBlank bool   `TargetBlank`
	Bz          string `Bz`
}

type Cate struct {
	Id  GUID   `bson:"_id,omitempty"`
	Uri string `Uri`
	Xx  `Xx`
}

func NewCate() *Cate {
	return &Cate{
		Id: mgo.NewGuid(),
		Xx: Xx{Lx: "cate"},
	}
}

func (p Cate) GetMc() string  { return p.Mc }
func (p Cate) GetUri() string { return p.Uri }
func (p Cate) GetQz() int     { return p.Qz }

type Info struct {
	Id      GUID              `bson:"_id,omitempty"`
	UserId  GUID              `UserId`
	Desc    string            `Desc`
	Zt      string            `Zt`
	PubDate time.Time         `PubDate`
	Fbdw    string            `Fbdw`
	Cates   []GUID            `Cates`
	Tags    []string          `Tags`
	Meta    map[string]string `Meta`
	Fj      []string          `Fj`
	Xx      `Xx`
}

func NewInfo(uid GUID, cid GUID) *Info {
	return &Info{
		Id:      mgo.NewGuid(),
		UserId:  uid,
		PubDate: time.Now(),
		Cates:   []GUID{cid},
		Xx:      Xx{Lx: "info"},
		Fj:      []string{},
		Tags:    []string{},
		Meta:    map[string]string{},
		Fbdw:    "",
	}
}

type SiteInfo struct {
}

type CmsCfg struct {
	Site SiteInfo `Site`
	Tags []string `Tags`
}
