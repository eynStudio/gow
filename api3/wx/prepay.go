package wx

import (
	"encoding/xml"
	"log"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/net/http"
)

type prePay struct {
	Error
}

func (p *prePay) Do(body, tradeNo, ip string, fee int) (payid string, err error) {
	var respBytes []byte
	var data string

	r := NewUnifyOrderReq(body, tradeNo, ip, fee)
	if data, p.Err = r.GetXml(); p.IsErr() {
		return
	}
	log.Println(data)
	respBytes, p.Err = http.PostWiHeader(unifiedorder, data, M{
		"Accept":       "application/xml",
		"Content-Type": "application/xml;charset=utf-8"})

	log.Println(string(respBytes))
	xmlResp := UnifyOrderResp{}
	if p.Err = xml.Unmarshal(respBytes, &xmlResp); p.IsErr() {
		return
	}

	if xmlResp.Return_code == "FAIL" {
		p.SetErr(xmlResp.Return_msg)
		return
	}
	p.LogErr()
	return xmlResp.Prepay_id, p.Err
}
