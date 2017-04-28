package cms

import (
	"errors"
	"log"
	"net/http"
	"time"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/db/filter"
	"github.com/eynstudio/gobreak/net/io"
	"github.com/eynstudio/gobreak/orm"
)

type CmsCtx struct {
	orm *orm.Orm
}

func (c *CmsCtx) SetOrm(orm *orm.Orm) { c.orm = orm }

func (c *CmsCtx) Orm() *orm.Orm {
	if c.orm != nil {
		return c.orm
	}
	return orm.GetOrmByName("gow.cms")
}

func (c *CmsCtx) CateTree() (tree *CateTree, err error) {
	var lst []CmsInfo
	s := c.Orm().Where(`json->>'IsCate'='true'`).Order("json->>'Ns'", "json->'Qz' desc").AllJson(&lst)
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
	c.Orm().WhereId(id).GetJson2(&m)
	return
}

func (c *CmsCtx) GetCateByUri(ns, mc string) (m CmsInfo) {
	c.Orm().Where(`json->>'Ns'=? and json->>'Mc'=?`, ns, mc).GetJson2(&m)
	return
}

func (c *CmsCtx) GetCateInfo(id GUID, f *filter.PageFilter) (m CateInfo) {
	c.Orm().WhereId(id).GetJson2(&m.Cate)
	p := c.Orm().Where(`json->'Cates' @> '"`+id.String()+`"'`).PageJson2(&m.Items, f)
	m.Total = p.Total
	return m
}

func (c *CmsCtx) SaveCate(m *CmsInfo) error {
	if m.Uid.IsEmpty() {
		return errors.New("NO UID")
	}
	return c.Orm().SaveJson(m.Id, m)
}

func (c *CmsCtx) GetInfo(id GUID) (m CmsInfo) {
	if id.IsEmpty() {
		m.Id = Guid()
		m.Lx = "info"
		m.Fbsj = time.Now()
		m.Cates = []GUID{Cate_Cgid}
		return
	}
	c.Orm().WhereId(id).GetJson2(&m)
	return
}

func (c *CmsCtx) SaveInfo(m *CmsInfo) error {
	log.Println(m)
	if m.Uid.IsEmpty() {
		return errors.New("NO UID")
	}
	return c.Orm().SaveJson(m.Id, m)
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
