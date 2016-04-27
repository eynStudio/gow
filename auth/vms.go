package auth

import (
	. "github.com/eynstudio/gobreak"
)

type Login struct {
	UserName string
	UserPwd  string
	AuthLx   string
}

type LoginOk struct {
	Id   string
	Name string
	Navs []*TreeNode
}

type LoginErr struct {
	Err string
}
