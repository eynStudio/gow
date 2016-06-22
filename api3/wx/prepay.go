package wx

import (
	"log"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/net/http"
)

type prePay struct {
	Error
}

func (p *prePay) Do(body, tradeNo, ip string, fee int) (payid string, err error) {
	var data string

	r := NewUnifyOrderReq(body, tradeNo, ip, fee)
	if data, p.Err = r.GetXml(); p.IsErr() {
		return
	}
	log.Println(data)

	xmlResp := UnifyOrderResp{}
	http.Post(unifiedorder, data, "").Header(M{
		"Accept":       "application/xml",
		"Content-Type": "application/xml;charset=utf-8"}).GetXml(&xmlResp)

	p.NoErrExec(func() {
		if xmlResp.Return_code == "FAIL" {
			p.SetErr(xmlResp.Return_msg)
		}
	})

	p.LogErr()
	return xmlResp.Prepay_id, p.Err
}
