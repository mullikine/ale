package builtin_test

import (
	"bytes"
	"testing"

	"github.com/kode4food/ale/core/bootstrap"
	"github.com/kode4food/ale/core/builtin"
	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/eval"
	"github.com/kode4food/ale/internal/assert"
	"github.com/kode4food/ale/namespace"
	"github.com/kode4food/ale/stdlib"
)

const stdoutName = "*out*"

func bindWrite(w stdlib.Writer) data.Call {
	return data.Call(func(args ...data.Value) data.Value {
		for _, af := range args {
			w.Write(af)
		}
		return nil
	})
}

func testOutput(t *testing.T, src, expected string) {
	as := assert.New(t)

	buf := bytes.NewBufferString("")
	w := stdlib.NewWriter(buf, stdlib.StrOutput)

	manager := namespace.NewManager()
	ns := manager.GetRoot()
	ns.Declare(stdoutName).Bind(data.Object{
		builtin.WriterKey: w,
		builtin.WriteKey:  bindWrite(w),
	})
	bootstrap.Into(manager)

	anon := manager.GetAnonymous()
	eval.String(anon, data.String(src))

	as.String(expected, buf.String())
}

func TestIOEval(t *testing.T) {
	testOutput(t, `
		(println "hello" "there")
	`, "hello there\n")

	testOutput(t, `
		(print "hello" "there")
	`, "hello there")

	testOutput(t, `
		(print "hello" 99)
	`, "hello 99")

	testOutput(t, `
		(prn "hello" "there")
	`, "\"hello\" \"there\"\n")

	testOutput(t, `
		(pr "hello" "there")
	`, "\"hello\" \"there\"")

	testOutput(t, `
		(pr "hello" 99)
	`, "\"hello\" 99")
}
