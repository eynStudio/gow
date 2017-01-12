package auth

import (
	"log"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gow/auth/group"
	"github.com/eynstudio/gow/auth/res"
	"github.com/eynstudio/gow/auth/role"
	"github.com/eynstudio/gow/auth/users"
	"github.com/eynstudio/gow/auth/x/redissess"

	"github.com/eynstudio/gobreak/orm"
	"github.com/eynstudio/gox/di"
)

func init() {
	Must(di.Reg(&AuthCtx{}))
}

type AuthCtx struct {
	*orm.Orm        `di:"*"`
	*group.GroupCtx `di:"*"`
	*role.RoleCtx   `di:"*"`
	*res.ResCtx     `di:"*"`
	*users.UserCtx  `di:"*"`
}

func (ac AuthCtx) GetGroupRoles(id GUID) (m GroupRoles) {
	m.Group, _ = ac.GroupCtx.Get(id)
	m.Roles, _ = ac.RoleCtx.AllAsTree()
	return
}

func (ac AuthCtx) GetRoleRes(id GUID) (m RoleRes) {
	m.Role, _ = ac.RoleCtx.Get(id)
	m.Res, _ = ac.ResCtx.AllAsTree()
	return
}

func (ac AuthCtx) GetOrgGroup(oid GUID) (m []GroupItem) {
	lst, _ := ac.GroupCtx.All(oid)
	for _, it := range lst {
		v := GroupItem{AuthGroup: it, Users: ac.UserCountByGroup(it.Id)}
		m = append(m, v)
	}
	return
}

func (ac AuthCtx) Login(req LoginReq) (resp LoginResp) {
	log.Println(req)

	//	if req.Mc == "pswang" {
	//		resp.Token = "eyn"
	//		resp.Ok()
	//		return
	//	}
	u, ok := ac.GetByMcPwd(req.Mc, SaltPwd(req.Pwd))
	log.Println(u, ok)
	if !ok || u.IsLock() {
		resp.ErrMsg("登录失败")
		return
	}

	resp.Token = Guid().String()
	//		resp.Id = u.Id
	redissess.SetSess(resp.Token, u.Id.String())
	resp.Ok()
	log.Println(resp)
	return
}

func (ac AuthCtx) GetUserNavs(uid GUID) Navs {
	lst := ac.GetUserRes(uid)
	tree := buildNavTree(lst)
	log.Println(tree)
	return tree.Navs
}

func (ac AuthCtx) GetUserRes(uid GUID) (lst []res.AuthRes) {
	sql := `json->'Id' in (
  SELECT jsonb_array_elements(role.json -> 'Res')->'ResId'
  FROM auth_role role
  WHERE role.json->>'Res' is not null  and (role.json -> 'Id' IN (
    SELECT jsonb_array_elements(g.json -> 'Roles')
    FROM auth_group g
    WHERE g.json -> 'Id' <@ (SELECT u.json -> 'Groups'
                             FROM auth_user u
                             WHERE id = ?)
  ))
) ORDER BY json->'Qz' DESC`
	err := ac.Orm.Where(sql, uid).AllJson(&lst).Err
	if err != nil {
		log.Println(err)
	}
	log.Println(lst)
	return lst
}
