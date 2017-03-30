package cms

import (
	"time"

	. "github.com/eynstudio/gobreak"
)

type CmsInfo struct {
	Id     GUID
	Uid    GUID   //发布者user.id
	Mc     string //名称，主标题
	Qz     int    //权重
	Ns     string //分类：树展示路径
	Lx     string //类型？//信息，分类，文件下载，链接，特殊模型（红头文件等）
	Ljdkfs bool   //链接打开方式，Target :_blank
	Zt     string //状态？
	Cjsj   time.Time
	Gxsj   time.Time
	Fbsj   time.Time `json:",omitempty"` //发布时间
	Nr     string    `json:",omitempty"` //内容
	Mc2    string    `json:",omitempty"` //副标题
	Ztp    string    `json:",omitempty"` //主图片，一般用于展示
	Fbdw   string    `json:",omitempty"` //发布单位（信息）
	Zy     string    `json:",omitempty"` //摘要
	Fj     []CmsFj   `json:",omitempty"` //附件列表
	Cates  []GUID    `json:",omitempty"` //分类列表（信息）
	Tags   []string  `json:",omitempty"`
	Bz     string    `json:",omitempty"` // 备注
	Ext    M         `json:",omitempty"`
	Fav    bool
}

func NewCate() *CmsInfo {
	return &CmsInfo{Id: Guid(), Lx: "list"}
}

func NewInfo(uid GUID, cid GUID) *CmsInfo {
	now := time.Now()
	return &CmsInfo{
		Id:    Guid(),
		Uid:   uid,
		Cjsj:  now,
		Gxsj:  now,
		Fbsj:  now,
		Cates: []GUID{cid},
		Lx:    "page",
	}
}

func (p CmsInfo) GetMc() string { return p.Mc }
func (p CmsInfo) GetNs() string { return p.Ns }
func (p CmsInfo) GetQz() int    { return p.Qz }

func (p CmsInfo) GetUri() string {
	if p.Ns == "" {
		return p.Mc
	}
	return p.Ns + `\` + p.Mc
}

type CmsFj struct {
	Mc  string
	Url string
}

type CmsCfg struct {
}
