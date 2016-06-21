package jpush

import (
	"encoding/base64"
	"errors"
	"strings"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/net/http"
)

const (
	JPUSH_VALIDATE = "https://api.jpush.cn/v3/push/validate"
	JPUSH_URL      = "https://api.jpush.cn/v3/push"
	BASE64         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var (
	appKey    string
	appSecret string
	auth      string
)

func Init(key, secret string) {
	appKey = key
	appSecret = secret
	auth = "Basic " + base64.NewEncoding(BASE64).EncodeToString([]byte(key+":"+secret))
}

func Send(push *JPush) (err error)     { return send(JPUSH_URL, push) }
func SendTest(push *JPush) (err error) { return send(JPUSH_VALIDATE, push) }

func send(url string, push *JPush) (err error) {
	data, err := http.PostJsonWiHeader(url, push, M{"Authorization": auth})
	if err != nil {
		return err
	}
	str := string(data)
	if strings.Contains(str, "msg_id") {
		return nil
	}
	return errors.New(str)
}

type JPush struct {
	Platform     interface{} `json:"platform"`
	Audience     interface{} `json:"audience"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Options      *Option     `json:"options,omitempty"`
}

func NewJPush() *JPush           { return &JPush{} }
func (p *JPush) Send() error     { return Send(p) }
func (p *JPush) SendTest() error { return SendTest(p) }

func (p *JPush) SetPlatform(all, android, ios, winphone bool) *JPush {
	if all {
		p.Platform = "all"
		p.Options = &Option{}
		p.Options.ApnsProduction = false
		return p
	}

	lst := make([]string, 0)
	if android {
		lst = append(lst, "android")
	}
	if ios {
		lst = append(lst, "ios")
		p.Options = &Option{}
		p.Options.ApnsProduction = false
	}
	if winphone {
		lst = append(lst, "winphone")
	}
	p.Platform = lst
	return p
}

func (p *JPush) SetNotification(alert string, android *AndroidNotice, ios *IOSNotice, win *WinPhoneNotice) *JPush {
	p.Notification = &Notice{Alert: alert, Android: android, IOS: ios, WINPhone: win}
	return p
}

func (p *JPush) SetMessage(content, title, content_type string, extras M) *JPush {
	p.Message = &Message{Content: content, Title: title, ContentType: content_type, Extras: extras}
	return p
}

func (p *JPush) SetAudience(all bool, tag, tag_and, alias, id []string) *JPush {
	if all {
		p.Audience = "all"
		return p
	}

	m := make(map[string][]string, 0)
	set := func(k string, v []string) {
		if v != nil {
			m[k] = v
		}
	}
	set("tag", tag)
	set("tag_and", tag_and)
	set("alias", alias)
	set("registration_id", id)
	p.Audience = m
	return p
}
