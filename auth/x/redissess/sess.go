package redissess

import (
	"github.com/eynstudio/gobreak"
	"github.com/eynstudio/gox/db/redisx"
	"github.com/eynstudio/gweb"
)

var (
	sess_user = "gweb:sess:user" //hash key sid uid
	user_sess = "gweb:user:sess" //hash key uid sid
)

func SetKeys(sessUser, userSess string) {
	sess_user, user_sess = sessUser, userSess
}

func Valid(c *gweb.Ctx) (httpCode int, s gobreak.IStatus, uid gobreak.GUID) {

	return
}

func HasSess(sid string) bool {
	n, err := redisx.Hexists(sess_user, sid)
	if err != nil || n == 0 {
		return false
	}
	return true
}

func GetSessUid(sid string) (uid string, err error) {
	return redisx.String(redisx.Hget(sess_user, sid))
}

func getSessSid(uid string) (sid string, err error) {
	return redisx.String(redisx.Hget(user_sess, uid))
}

func SetSess(sid, uid string) {
	old_sid, _ := getSessSid(uid)
	redisx.Hdel(sess_user, old_sid)
	redisx.Hset(user_sess, uid, sid)
	redisx.Hset(sess_user, sid, uid)
}

func DelSess(sid string) error {
	uid, _ := GetSessUid(sid)
	redisx.Hdel(user_sess, uid)
	redisx.Hdel(sess_user, sid)
	return nil
}
