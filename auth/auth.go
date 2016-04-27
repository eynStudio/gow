package auth

import (
	"golang.org/x/crypto/scrypt"
)

var (
	salt string
)

func SetSalt(s string) { salt = s }

func SaltPwd(pwd string) string {
	for i := 0; i < 10; i++ {
		if k, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32); err == nil {
			return string(k)
		}
	}
	return salt + pwd
}
