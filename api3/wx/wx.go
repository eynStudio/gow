package wx

const (
	unifiedorder = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	orderquery   = "https://api.mch.weixin.qq.com/pay/orderquery"
)

var (
	wxAppid     string
	wxMchid     string
	wxApikey    string
	wxNotifyUrl string
)

func Init(appid, mchid, apikey, notifyurl string) {
	wxApikey = apikey
	wxAppid = appid
	wxMchid = mchid
	wxNotifyUrl = notifyurl
}

func PrePay(body, tradeNo, ip string, fee int) (payreq *PayReq, err error) {
	p := &prePay{}
	return p.Do(body, tradeNo, ip, fee)
}

func OrderQuery(tradeNo string) (resp OrderQueryResp, err error) {
	p := &orderQuery{}
	return p.Do(tradeNo)
}
