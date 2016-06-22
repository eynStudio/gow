package jpush

import . "github.com/eynstudio/gobreak"

type Notice struct {
	Alert    string          `json:"alert,omitempty"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
}

func NewNotice(alert string, ext M) *Notice {
	a := &AndroidNotice{Alert: alert, Extras: ext}
	b := &IOSNotice{Alert: alert, Extras: ext}
	c := &WinPhoneNotice{Alert: alert, Extras: ext}
	return &Notice{Alert: "", Android: a, IOS: b, WINPhone: c}
}

type AndroidNotice struct {
	Alert     string `json:"alert"`
	Title     string `json:"title,omitempty"`
	BuilderId int    `json:"builder_id,omitempty"`
	Extras    M      `json:"extras,omitempty"`
}

type IOSNotice struct {
	Alert            string `json:"alert"`
	Sound            string `json:"sound,omitempty"`
	Badge            int    `json:"badge,omitempty"`
	ContentAvailable bool   `json:"Content-available,omitempty"`
	Category         string `json:"category,omitempty"`
	Extras           M      `json:"extras,omitempty"`
}

type WinPhoneNotice struct {
	Alert    string `json:"alert"`
	Title    string `json:"title,omitempty"`
	OpenPage string `json:"_open_page,omitempty"`
	Extras   M      `json:"extras,omitempty"`
}

type Option struct {
	SendNo          int   `json:"sendno,omitempty"`
	TimeLive        int   `json:"time_to_live,omitempty"`
	ApnsProduction  bool  `json:"apns_production"`
	OverrideMsgId   int64 `json:"override_msg_id,omitempty"`
	BigPushDuration int   `json:"big_push_duration,omitempty"`
}

type Message struct {
	Content     string `json:"msg_content"`
	Title       string `json:"title,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Extras      M      `json:"extras,omitempty"`
}
