package auth_test

import (
	"testing"

	"github.com/eynstudio/gobreak/di2"
	"github.com/eynstudio/gow/auth"
	_ "github.com/eynstudio/gow/auth/pg"
)

type A struct {
	auth.IAuthCtx `di:"*"`
}

func TestCtx(t *testing.T) {
	var a A
	di2.Reg(&a)
	a.Hi()
}
