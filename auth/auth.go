package auth

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

var (
	salt string
)

func SetSalt(s string) { salt = s }

func SaltPwd(pwd string) string {
	for i := 0; i < 10; i++ {
		if k, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32); err == nil {
			m5 := md5.New()
			m5.Write(k)
			return hex.EncodeToString(m5.Sum(nil))
			//			return string(k)
		}
	}
	return salt + pwd
}

func Temp(pwd string) string {
	m5 := md5.New()
	m5.Write([]byte(pwd))
	return hex.EncodeToString(m5.Sum(nil))
}
