package wx

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"sort"
)

// Sign 微信支付签名.
func Sign(params map[string]string) string {

	h := md5.New()
	bufw := bufio.NewWriterSize(h, 128)

	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		bufw.WriteString(k)
		bufw.WriteByte('=')
		bufw.WriteString(v)
		bufw.WriteByte('&')
	}
	bufw.WriteString("key=")
	bufw.WriteString(wxApikey)

	bufw.Flush()
	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}
