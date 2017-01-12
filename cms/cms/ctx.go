package cms

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
	"github.com/eynstudio/gox/utils"
)

func init() {
	Must(di.Reg(&CmsCtx{}))
}

type CmsCtx struct {
	*orm.Orm `di:"*"`
}

func (c *CmsCtx) CateTree() (tree utils.TreeNodes, err error) {
	var lst []CmsCate
	if err = c.Orm.AllJson(&lst); err != nil {
		return nil, err
	}
	return utils.BuildTree(lst), nil
}
