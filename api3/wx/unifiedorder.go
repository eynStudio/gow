package wx

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type UnifyOrderReq struct {
	Appid            string `xml:"appid"`
	Mch_id           string `xml:"mch_id"`
	Nonce_str        string `xml:"nonce_str"`        //随机字符串，不长于32位。
	Sign             string `xml:"sign"`             //签名
	Body             string `xml:"body"`             //商品描述
	Out_trade_no     string `xml:"out_trade_no"`     //商户订单号
	Total_fee        int    `xml:"total_fee"`        // 订单总金额，单位为分
	Spbill_create_ip string `xml:"spbill_create_ip"` //用户端实际ip
	Notify_url       string `xml:"notify_url"`       //接收微信支付异步通知回调地址
	Trade_type       string `xml:"trade_type"`       //支付类型APP
}

func NewUnifyOrderReq(body, tradeNo, ip string, fee int) *UnifyOrderReq {
	m := &UnifyOrderReq{}
	m.Appid = wxAppid
	m.Mch_id = wxMchid
	m.Nonce_str = tradeNo
	m.Body = body
	m.Out_trade_no = tradeNo
	m.Total_fee = fee
	m.Spbill_create_ip = ip
	m.Notify_url = wxNotifyUrl
	m.Trade_type = "APP"
	mp := m.getSignMap()
	m.Sign = Sign(mp)
	return m
}

func (p UnifyOrderReq) getSignMap() (m map[string]string) {
	m = make(map[string]string, 0)
	m["appid"] = p.Appid
	m["body"] = p.Body
	m["mch_id"] = p.Mch_id
	m["notify_url"] = p.Notify_url
	m["trade_type"] = p.Trade_type
	m["spbill_create_ip"] = p.Spbill_create_ip
	m["total_fee"] = strconv.Itoa(p.Total_fee)
	m["out_trade_no"] = p.Out_trade_no
	m["nonce_str"] = p.Nonce_str
	return
}

func (p UnifyOrderReq) GetXml() (string, error) {
	data, err := xml.Marshal(p)
	if err != nil {
		return "", err
	}

	str := strings.Replace(string(data), "UnifyOrderReq", "xml", -1)
	return str, nil
}

type UnifyOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
}
