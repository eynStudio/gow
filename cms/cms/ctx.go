package cms

import (
	"errors"
	"log"
	"time"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

func init() {
	Must(di.Reg(&CmsCtx{}))
}

type CmsCtx struct {
	*orm.Orm `di:"*"`
}

func (c *CmsCtx) CateTree() (tree *CateTree, err error) {
	var lst []CmsInfo
	s := c.Orm.Order("json->>'Ns'", "json->'Qz'").AllJson(&lst)
	if s.IsErr() {
		return nil, s.Err
	}
	for _, it := range lst {
		log.Printf("%#v", it)
		log.Println(it.Ns, it.Qz, it.Mc, it.GetUri())
	}
	tree = NewCateTree()
	log.Println(lst)
	tree.Build(lst)
	return tree, nil
}

func (c *CmsCtx) GetCate(id GUID) (m CmsInfo) {
	if id.IsEmpty() {
		m.Id = Guid()
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
		m.Cates = make([]GUID, 0)
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
