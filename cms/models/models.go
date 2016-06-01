package models

import (
	"time"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/db/mgo"
)

type Xx struct {
	Id      GUID      `bson:"_id"` //ID
	Mc      string    `Mc`         //名称，主标题
	Mc2     string    `Mc2`        //副标题
	Qz      int       `Qz`         //权重
	Lx      string    `Lx`         //类型？
	Nr      string    `Nr`         //内容
	Ztp     string    `Ztp`        //主图片，一般用于展示
	Ljdkfs  bool      `Ljdkfs`     //链接打开方式，Target :_blank
	Bz      string    `Bz`         // 备注
	Created time.Time `Created`
	Updated time.Time `Updated`
	Ext     M         `Ext`
}

type Cate struct {
	Xx  `bson:",inline"`
	Uri string `Uri` //树展示路径
}

func NewCate() *Cate {
	return &Cate{Xx: Xx{Id: mgo.NewGuid(), Lx: "cate"}}
}

func (p Cate) GetMc() string  { return p.Mc }
func (p Cate) GetUri() string { return p.Uri }
func (p Cate) GetQz() int     { return p.Qz }

type Info struct {
	Xx    `bson:",inline"`
	Uid   GUID      `Uid`  //发布者user.id
	Zy    string    `Zy`   //摘要
	Zt    string    `Zt`   //状态？
	Fbsj  time.Time `Fbsj` //发布时间
	Fbdw  string    `Fbdw` //发布单位
	Fj    []string  `Fj`   //附件列表
	Cates []GUID    `Cates`
	Tags  []string  `Tags`
}

func NewInfo(uid GUID, cid GUID) *Info {
	return &Info{
		Uid:   uid,
		Fbsj:  time.Now(),
		Cates: []GUID{cid},
		Xx:    Xx{Id: mgo.NewGuid(), Ext: make(M, 0), Lx: "info"},
		Fj:    []string{},
		Tags:  []string{},
	}
}

type CmsCfg struct {
}
