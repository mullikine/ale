package bootstrap_test

import (
	"testing"

	"gitlab.com/kode4food/ale/compiler/encoder"
	"gitlab.com/kode4food/ale/core/bootstrap"
	"gitlab.com/kode4food/ale/data"
	"gitlab.com/kode4food/ale/internal/assert"
	"gitlab.com/kode4food/ale/stdlib"
)

func TestDevNullManager(t *testing.T) {
	as := assert.New(t)

	manager := bootstrap.DevNullManager()
	ns := manager.GetRoot()

	_, ok := ns.Resolve("*args*")
	as.False(ok)

	e, ok := ns.Resolve("*in*")
	as.True(ok && e.IsBound())
	r, ok := e.Value().(stdlib.Reader)
	as.True(ok)
	as.True(r.IsEmpty())
}

func TestTopLevelManager(t *testing.T) {
	as := assert.New(t)

	manager := bootstrap.TopLevelManager()
	ns := manager.GetRoot()

	e, ok := ns.Resolve("*args*")
	as.True(ok && e.IsBound())

	_, ok = e.Value().(data.Vector)
	as.True(ok)
}

func TestBootstrapInto(t *testing.T) {
	as := assert.New(t)

	manager := bootstrap.TopLevelManager()
	bootstrap.Into(manager)
	ns := manager.GetRoot()

	e, ok := ns.Resolve("def")
	as.True(ok && e.IsBound())

	_, ok = e.Value().(encoder.Call)
	as.True(ok)
}
