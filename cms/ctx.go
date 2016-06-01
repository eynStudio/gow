package cms

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
	"github.com/eynstudio/gobreak/di"
	"github.com/eynstudio/gow/cms/repo"
	"gopkg.in/mgo.v2/bson"
)

func Init(domainRepo DomainRepo, aggCmdHandler AggCmdHandler, eventBus EventBus) {
	repoCate := NewCateRepo()
	di.MapAs(repoCate, (*repo.ICateRepo)(nil)).Apply(repoCate.MgoRepo)

	repoInfo := NewInfoRepo()
	di.MapAs(repoInfo, (*repo.IInfoRepo)(nil)).Apply(repoInfo.MgoRepo)

	di.Root.ApplyAndMap(&CmsCtx{})
}

type CmsCtx struct {
	repo.ICateRepo `di`
	repo.IInfoRepo `di`
}

func (p *CmsCtx) GetCateTree() []*TreeNode {
	return BuildTree(p.ICateRepo.All())
}

func (p *CmsCtx) GetInfos(cid GUID) []T {
	return p.IInfoRepo.Find(bson.M{"Cates": cid})
}
