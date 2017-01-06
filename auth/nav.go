package auth

import (
	"strings"

	"github.com/ahmetalpbalkan/go-linq"
	"github.com/eynstudio/gow/auth/res"
)

type Nav struct {
	Mc   string
	To   string
	Navs `json:",omitempty"`
	Qz   int
}

func getTo(str string) string { return "/" + strings.Replace(str, ".", "/", -1) }

type Navs []*Nav

func buildNavTree(src interface{}) *Nav {
	root := &Nav{Mc: "", To: "", Navs: make(Navs, 0)}
	buildNavNodes(src, root, "")
	return root
}

func buildNavNodes(src interface{}, r *Nav, prefix string) {
	results := queryChildren(src, prefix)
	for _, it := range results {
		x := it.(res.AuthRes)
		child := &Nav{Mc: x.Mc, To: getTo(x.Ns), Qz: x.Qz, Navs: make(Navs, 0)}
		r.Navs = append(r.Navs, child)
		buildNavNodes(src, child, x.Ns+".")
	}
}

func queryChildren(ss interface{}, prefix string) []interface{} {
	return linq.From(ss).Where(func(s interface{}) bool {
		x := s.(res.AuthRes)
		last := strings.TrimPrefix(x.Ns, prefix)
		return strings.HasPrefix(x.Ns, prefix) && !strings.Contains(last, ".")
	}).OrderByDescending(func(a interface{}) interface{} {
		return a.(res.AuthRes).Qz
	}).Results()
}
