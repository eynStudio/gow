package wx

import (
	"fmt"
	"time"

	. "github.com/eynstudio/gobreak"
)

type PayReq struct {
	Appid     string
	Partnerid string
	Prepayid  string
	Package   string
	Noncestr  string
	Timestamp string
	Sign      string
}

func NewPayReq(preid string) *PayReq {
	m := &PayReq{}
	m.Appid = wxAppid
	m.Partnerid = wxMchid
	m.Prepayid = preid
	m.Package = "Sign=WXPay"
	m.Noncestr = Uuid4().String()[:32]
	m.Timestamp = fmt.Sprint(time.Now().Unix())
	mp := m.getSignMap()
	m.Sign = Sign(mp)
	return m
}

func (p *PayReq) getSignMap() (m map[string]string) {
	m = make(map[string]string, 0)
	m["appid"] = p.Appid
	m["partnerid"] = p.Partnerid
	m["prepayid"] = p.Prepayid
	m["package"] = p.Package
	m["noncestr"] = p.Noncestr
	m["timestamp"] = p.Timestamp
	return
}
