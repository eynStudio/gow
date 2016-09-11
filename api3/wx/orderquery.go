package wx

import (
	"encoding/xml"
	"strings"

	"github.com/eynstudio/gobreak/net/http"

	. "github.com/eynstudio/gobreak"
)

type orderQuery struct {
	Error
}

func (p *orderQuery) Do(tradeNo string) (xmlResp OrderQueryResp, err error) {
	var data string

	r := NewOrderQueryReq(tradeNo)
	if data, p.Err = r.GetXml(); p.IsErr() {
		return
	}
	http.Post(orderquery, data, "").Header(M{
		"Accept":       "application/xml",
		"Content-Type": "application/xml;charset=utf-8"}).GetXml(&xmlResp)

	p.NoErrExec(func() {
		if xmlResp.Return_code == "FAIL" {
			p.SetErr(xmlResp.Return_msg)
		}
	})

	p.LogErr()
	return xmlResp, p.Err
}

type OrderQueryReq struct {
	Appid        string `xml:"appid"`
	Mch_id       string `xml:"mch_id"`
	Nonce_str    string `xml:"nonce_str"`    //随机字符串，不长于32位。
	Sign         string `xml:"sign"`         //签名
	Out_trade_no string `xml:"out_trade_no"` //商户订单号
}

func NewOrderQueryReq(tradeNo string) *OrderQueryReq {
	m := &OrderQueryReq{}
	m.Appid = wxAppid
	m.Mch_id = wxMchid
	m.Nonce_str = tradeNo
	m.Out_trade_no = tradeNo
	mp := m.getSignMap()
	m.Sign = Sign(mp)
	return m
}

func (p OrderQueryReq) getSignMap() (m map[string]string) {
	m = make(map[string]string, 0)
	m["appid"] = p.Appid
	m["mch_id"] = p.Mch_id
	m["out_trade_no"] = p.Out_trade_no
	m["nonce_str"] = p.Nonce_str
	return
}

func (p OrderQueryReq) GetXml() (string, error) {
	data, err := xml.Marshal(p)
	if err != nil {
		return "", err
	}

	str := strings.Replace(string(data), "OrderQueryReq", "xml", -1)
	return str, nil
}

type OrderQueryResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Trade_state string `xml:"trade_state"`
	Total_fee   int    `xml:"total_fee"`
}
