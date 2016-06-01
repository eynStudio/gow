package cms

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/ddd"
	"github.com/eynstudio/gobreak/di"
	"gopkg.in/mgo.v2/bson"
)

func Init(domainRepo DomainRepo, aggCmdHandler AggCmdHandler, eventBus EventBus) {
	repoCate := NewCateRepo()
	di.Root.Map(repoCate).Apply(repoCate.MgoRepo)

	repoInfo := NewInfoRepo()
	di.Root.Map(repoInfo).Apply(repoInfo.MgoRepo)

	di.Root.ApplyAndMap(&CmsCtx{})
}

type CmsCtx struct {
	*CateRepo `di`
	*InfoRepo `di`
}

func (p *CmsCtx) GetCateTree() []*TreeNode {
	return BuildTree(p.CateRepo.All())
}

func (p *CmsCtx) GetInfos(cid GUID) []T {
	return p.InfoRepo.Find(bson.M{"Cates": cid})
}
