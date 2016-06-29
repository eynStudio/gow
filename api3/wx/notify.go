package wx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var OkNotifyResp = NotifyResp{Return_code: "SUCCESS", Return_msg: "OK"}
var FaildSignVerify = NotifyResp{Return_code: "FAIL", Return_msg: "failed to verify sign"}

type NotifyResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
}

func (p *NotifyResp) Ok() {
	p.Return_code = "SUCCESS"
	p.Return_msg = "OK"
}

func (p *NotifyResp) Fail(msg string) {
	p.Return_code = "FAIL"
	p.Return_msg = msg
}

func (p *NotifyResp) GetXml() string {
	bytes, _ := xml.Marshal(p)
	return strings.Replace(string(bytes), "NotifyResp", "xml", -1)
}

type NotifyReq struct {
	Return_code    string `xml:"return_code"`
	Return_msg     string `xml:"return_msg"`
	Appid          string `xml:"appid"`
	Mch_id         string `xml:"mch_id"`
	Nonce          string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	Result_code    string `xml:"result_code"`
	Openid         string `xml:"openid"`
	Is_subscribe   string `xml:"is_subscribe"`
	Trade_type     string `xml:"trade_type"`
	Bank_type      string `xml:"bank_type"`
	Total_fee      int    `xml:"total_fee"`
	Fee_type       string `xml:"fee_type"`
	Cash_fee       int    `xml:"cash_fee"`
	Cash_fee_Type  string `xml:"cash_fee_type"`
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	Attach         string `xml:"attach"`
	Time_end       string `xml:"time_end"`
}

func (p *NotifyReq) getNotifyMap() (m map[string]string) {
	m = make(map[string]string, 0)
	m["return_code"] = p.Return_code
	m["return_msg"] = p.Return_msg
	m["appid"] = p.Appid
	m["mch_id"] = p.Mch_id
	m["nonce_str"] = p.Nonce
	m["result_code"] = p.Result_code
	m["openid"] = p.Openid
	m["is_subscribe"] = p.Is_subscribe
	m["trade_type"] = p.Trade_type
	m["bank_type"] = p.Bank_type
	m["total_fee"] = strconv.Itoa(p.Total_fee)
	m["fee_type"] = p.Fee_type
	m["cash_fee"] = strconv.Itoa(p.Cash_fee)
	m["cash_fee_type"] = p.Cash_fee_Type
	m["transaction_id"] = p.Transaction_id
	m["out_trade_no"] = p.Out_trade_no
	m["attach"] = p.Attach
	m["time_end"] = p.Time_end
	return
}

func CheckNotify(r *http.Request) (n *NotifyReq, success bool, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取http body失败，原因!", err)
		return nil, false, err
	}
	defer r.Body.Close()

	var mr NotifyReq
	fmt.Println("微信支付异步通知，HTTP Body:", string(body))
	err = xml.Unmarshal(body, &mr)
	if err != nil {
		fmt.Println("解析HTTP Body格式到xml失败，原因!", err)
		return nil, false, err
	}

	sign := Sign(mr.getNotifyMap())
	fmt.Println("计算出来的sign:", sign)
	fmt.Println("微信异步通知sign: ", mr.Sign)

	return &mr, sign == mr.Sign, nil
}
