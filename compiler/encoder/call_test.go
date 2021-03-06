package encoder_test

import (
	"testing"

	"github.com/kode4food/ale/compiler/encoder"
	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/internal/assert"
)

func TestCall(t *testing.T) {
	as := assert.New(t)
	f1 := func(_ encoder.Type, _ ...data.Value) {}
	c1 := encoder.Call(f1)
	as.String("encoder-call", c1.Type())
	as.Contains(`:type encoder-call`, c1)
	as.Contains(`:instance `, c1)
}
