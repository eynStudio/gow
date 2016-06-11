package auth

import (
	"github.com/eynstudio/gobreak/db"
	"github.com/eynstudio/gobreak/db/filter"
	. "github.com/eynstudio/gow/auth/models"

	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/gobreak/db/mgo"
	"github.com/eynstudio/gow/auth/models"
	"github.com/eynstudio/gow/auth/repo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var _ repo.IUserRepo = new(UserRepo)

type UserRepo struct {
	MgoRepo
}

func NewUserRepo() *UserRepo {
	return &UserRepo{NewMgoRepo("AuthUser", func() T { return &User{} })}
}

func (p *UserRepo) GetById(id GUID) (m *models.User, err error) {
	m = new(models.User)
	p.GetAs(id, &m)
	if m.Id == "" {
		err = db.DbNotFound
	}
	return
}

func (p *UserRepo) HasUserMc(mc string) (has bool, err error) {
	sess := p.CopySession()
	defer sess.Close()
	var n int
	n, err = p.C(sess).Find(bson.M{"Mc": mc}).Count()
	return n > 0, err
}
func (p *UserRepo) CheckPwd(id GUID, pwd string) (has bool, err error) {
	var n int
	p.Sess(func(c *mgo.Collection) {
		n, err = c.Find(bson.M{"_id": id, "Pwd": pwd}).Count()
	})
	return n > 0, err
}
func (p *UserRepo) GetUserByMc(mc string) (u *User, ok bool) {
	sess := p.CopySession()
	defer sess.Close()
	u = new(User)
	err := p.C(sess).Find(bson.M{"Mc": mc}).One(u)
	return u, err == nil
}

func (p *UserRepo) GetUserByMcPwd(name, pwd string) (u *User, ok bool) {
	sess := p.CopySession()
	defer sess.Close()
	u = new(User)
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

func (p *UserRepo) UpdateNc(uid GUID, nc string)   { p.UpdateSetFiled(uid, "Nc", nc) }
func (p *UserRepo) UpdatePwd(uid GUID, pwd string) { p.UpdateSetFiled(uid, "Pwd", pwd) }
func (p *UserRepo) UpdateImg(uid GUID, img string) { p.UpdateSetFiled(uid, "Img", img) }
