package cms

import (
	"errors"
	"log"
	"net/http"
	"time"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/net/io"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

const (
	Cate_Cgid = GUID("7cae5572-02ef-11e7-902f-c07cd130ee8a")
)

func init() {
	Must(di.Reg(&CmsCtx{}))
}

type CmsCtx struct {
	*orm.Orm `di:"*"`
}

func (c *CmsCtx) CateTree() (tree *CateTree, err error) {
	var lst []CmsInfo
	s := c.Orm.Where(`json->>'IsCate'='true'`).Order("json->>'Ns'", "json->'Qz' desc").AllJson(&lst)
	if s.IsErr() {
		return nil, s.Err
	}
	tree = NewCateTree()
	log.Println(lst)
	tree.Build(lst)
	return tree, nil
}

func (c *CmsCtx) GetCate(id GUID) (m CmsInfo) {
	if id.IsEmpty() {
		m.Id = Guid()
		m.IsCate = true
		return
	}
	c.Orm.WhereId(id).GetJson2(&m)
	return
}

func (c *CmsCtx) GetCateInfo(id GUID) (m CateInfo) {
	c.Orm.WhereId(id).GetJson2(&m.Cate)
	s := c.Orm.Where(`json->'Cates' @> '"` + id.String() + `"'`).AllJson(&m.Items)
	s.LogErr()
	return
}

func (c *CmsCtx) SaveCate(m *CmsInfo) error {
	if m.Uid.IsEmpty() {
		return errors.New("NO UID")
	}
	return c.Orm.SaveJson(m.Id, m)
}

func (c *CmsCtx) GetInfo(id GUID) (m CmsInfo) {
	if id.IsEmpty() {
		m.Id = Guid()
		m.Lx = "info"
		m.Fbsj = time.Now()
		m.Cates = []GUID{Cate_Cgid}
		return
	}
	c.Orm.WhereId(id).GetJson2(&m)
	return
}

func (c *CmsCtx) SaveInfo(m *CmsInfo) error {
	if m.Uid.IsEmpty() {
		return errors.New("NO UID")
	}
	return c.Orm.SaveJson(m.Id, m)
}

func UploadImg(req *http.Request) io.UrlStatus {
	saveFile := io.NewSaveYyyyMmFile("./files", Guid().String(), []string{".png", ".jpg", ".jpeg", ".gif", ".bmp"})
	saveFile.Save(req)
	return saveFile.GetUrlStatus()
}

func UploadFile(req *http.Request) (m CmsFj) {
	saveFile := io.NewSaveYyyyMmFile("./files", Guid().String(), nil)
	saveFile.Save(req)
	m.Mc = saveFile.GetFileName()
	m.Url = saveFile.GetUrlStatus().Url
	return
}
