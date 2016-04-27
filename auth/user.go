package auth

import (
	. "github.com/eynstudio/gow/auth/user"

	"github.com/eynstudio/gobreak/db"
	"github.com/eynstudio/gobreak/db/filter"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	"gopkg.in/mgo.v2/bson"
)

type UserRepo struct {
	MgoRepo
}

func NewUserRepo() *UserRepo {
	return &UserRepo{NewMgoRepo("AuthUser", func() T { return &User{} })}
}

func (p *UserRepo) GetById(id GUID) (u User, ok bool) {
	p.GetAs(id, &u)
	return u, u.Id != ""
}

func (p *UserRepo) HasUserByMc(mc string) (has bool, err error) {
	sess := p.CopySession()
	defer sess.Close()
	var n int
	n, err = p.C(sess).Find(bson.M{"Mc": mc}).Count()
	return n > 0, err
}

func (p *UserRepo) GetUserByMc(mc string) (u User, ok bool) {
	sess := p.CopySession()
	defer sess.Close()
	err := p.C(sess).Find(bson.M{"Mc": mc}).One(&u)
	return u, err == nil
}

func (p *UserRepo) GetUser(name, pwd string) (u User, ok bool) {
	sess := p.CopySession()
	defer sess.Close()
	err := p.C(sess).Find(bson.M{"Mc": name, "Pwd": pwd}).One(&u)
	return u, err == nil
}

func (p *UserRepo) GetUserByLx(name, lx string) (u User, ok bool) {
	sess := p.CopySession()
	defer sess.Close()
	err := p.C(sess).Find(bson.M{"Auth": bson.M{"Mc": name, "Lx": lx}}).One(&u)
	return u, err == nil
}

func (p *UserRepo) ByGroup(gid GUID, pf *filter.PageFilter) (pager db.Paging) {
	sess := p.CopySession()
	defer sess.Close()
	return p.Page(pf, bson.M{"Groups": gid})
}

func (p *UserRepo) CountByGroup(gid GUID) int {
	sess := p.CopySession()
	defer sess.Close()
	n, _ := p.C(sess).Find(bson.M{"Groups": gid}).Count()
	return n
}

func (p *UserRepo) AddGroup(uid, gid GUID) {
	p.Save(uid, bson.M{"$addToSet": bson.M{"Groups": gid}})
}
func (p *UserRepo) DelGroup(uid, gid GUID) {
	p.Save(uid, bson.M{"$pull": bson.M{"Groups": gid}})
}
