package dayu

import (
	"crypto/md5"
	"errors"
	"fmt"
	. "github.com/eynstudio/gobreak"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Alidayu_URL        string = "http://gw.api.taobao.com/router/rest"
	Alidayu_SendSMS    string = "alibaba.aliqin.fc.sms.num.send"
	Alidayu_CallTTS    string = "alibaba.aliqin.fc.tts.num.singlecall"
	Alidayu_CallVoice  string = "alibaba.aliqin.fc.voice.num.singlecall"
	Alidayu_CallDouble string = "alibaba.aliqin.fc.voice.num.doublecall"
)

var (
	ErrRequired   error = errors.New("AppKey or AppSecret is required!")
	ErrParameters error = errors.New("Parameters are not complete!")
	AppKey        string
	AppSecret     string
)

func SendSMS(rec_num, sms_free_sign_name, sms_template_code, sms_param string) (err error) {
	if rec_num == "" || sms_free_sign_name == "" || sms_template_code == "" {
		return ErrParameters
	}

	m := getCommonCfg()
	m["method"] = Alidayu_SendSMS
	m["sms_type"] = "normal"
	m["sms_free_sign_name"] = sms_free_sign_name
	m["rec_num"] = rec_num
	m["sms_template_code"] = sms_template_code
	m["sms_param"] = sms_param
	return post(m)
}

func post(m M) (err error) {
	if AppKey == "" || AppSecret == "" {
		return ErrRequired
	}

	signedData := getSignedData(m)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", Alidayu_URL, strings.NewReader(signedData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = int64(len(signedData))

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	response := string(data)
	if strings.Contains(response, "success") {
		return nil
	} else {
		return errors.New(response)
	}
}

func getCommonCfg() (m M) {
	m = make(M)
	m["app_key"] = AppKey
	m["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	m["format"] = "json"
	m["v"] = "2.0"
	m["sign_method"] = "md5"
	return m
}

func getSignedData(m M) string {
	v := url.Values{}

	keys := m.GetSortedKeys()
	signStr := AppSecret
	for _, k := range keys {
		v.Set(k, m.GetStr(k))
		signStr += k + m.GetStr(k)
	}
	signStr += AppSecret

	v.Set("sign", strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(signStr)))))
	return v.Encode()
}
