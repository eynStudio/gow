package alidayu

type DayuResp struct {
	SmsResp   `json:"alibaba_aliqin_fc_sms_num_send_response"`
	ErrorResp `json:"error_response"`
}

type SmsResp struct {
	Result `json:"result"`
}

type Result struct {
	ErrCode string `json:"err_code"`
	Model   string
	Success bool
}

type ErrorResp struct {
	Code    int
	Msg     string
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}
