package cms

import "github.com/eynstudio/gobreak"

type CateInfo struct {
	Cate  CmsInfo
	Items []CmsInfo
	Total int
}

type CateNode struct {
	Id    gobreak.GUID
	Mc    string
	Ns    string
	Uri   string
	Nodes []*CateNode
}

type CateTree struct {
	uriMap map[string]*CateNode
	Nodes  []*CateNode
}

func NewCateTree() *CateTree {
	return &CateTree{
		Nodes:  make([]*CateNode, 0),
		uriMap: make(map[string]*CateNode, 0),
	}
}
func (ct *CateTree) Build(all []CmsInfo) {
	for _, it := range all {
		n := newCateNode(it)
		ct.uriMap[n.Uri] = n
		if n.Ns == "" {
			ct.Nodes = append(ct.Nodes, n)
		} else {
			if _, ok := ct.uriMap[n.Ns]; ok {
				ct.uriMap[n.Ns].Nodes = append(ct.uriMap[n.Ns].Nodes, n)
			} else {
				ct.Nodes = append(ct.Nodes, n)
			}
		}
	}
}

func newCateNode(m CmsInfo) *CateNode {
	return &CateNode{
		Id:    m.Id,
		Mc:    m.Mc,
		Ns:    m.Ns,
		Uri:   m.GetUri(),
		Nodes: make([]*CateNode, 0),
	}
}
