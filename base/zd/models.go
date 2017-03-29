package zd

import (
	. "github.com/eynstudio/gobreak"
)

type XtZd struct {
	Id    GUID
	Mc    string
	Uri   string
	Qz    int
	Bz    string
	Items []ZdItem
}

func (p XtZd) GetMc() string  { return p.Mc }
func (p XtZd) GetUri() string { return p.Uri }
func (p XtZd) GetQz() int     { return p.Qz }

type ZdItem struct {
	Dm string
	Mc string
	Jc string
	Qz int //小于0，表示不可见
	Bz string
}

func (p XtZd) GetItemByDm(dm string) (ZdItem, bool) {
	for _, it := range p.Items {
		if it.Dm == dm {
			return it, true
		}
	}
	return ZdItem{}, false
}

func (p XtZd) GetItemByMc(mc string) (ZdItem, bool) {
	for _, it := range p.Items {
		if it.Mc == mc {
			return it, true
		}
	}
	return ZdItem{}, false
}

func (p XtZd) GetItemByJc(jc string) (ZdItem, bool) {
	for _, it := range p.Items {
		if it.Jc == jc {
			return it, true
		}
	}
	return ZdItem{}, false
}

func (p XtZd) GetDmByMc(mc string) string {
	it, _ := p.GetItemByMc(mc)
	return it.Dm
}

func (p XtZd) GetMcByDm(dm string) string {
	it, _ := p.GetItemByDm(dm)
	return it.Mc
}

func (p XtZd) GetJcByDm(dm string) string {
	it, _ := p.GetItemByDm(dm)
	return it.Jc
}

func (p XtZd) GetDmByJc(jc string) string {
	it, _ := p.GetItemByJc(jc)
	return it.Dm
}
