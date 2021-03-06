package builtin

import (
	"bytes"

	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/stdlib"
)

const emptyString = data.String("")

// Str converts the provided arguments to an undelimited string
func Str(args ...data.Value) data.Value {
	v := data.NewVector(args...)
	return stdlib.SequenceToStr(v)
}

// ReaderStr converts the provided arguments to a delimited string
func ReaderStr(args ...data.Value) data.Value {
	if len(args) == 0 {
		return emptyString
	}

	var b bytes.Buffer
	b.WriteString(data.MaybeQuoteString(args[0]))
	for _, f := range args[1:] {
		b.WriteString(" ")
		b.WriteString(data.MaybeQuoteString(f))
	}
	return data.String(b.String())
}

// IsString returns whether or not the provided value is a string
func IsString(args ...data.Value) data.Value {
	_, ok := args[0].(data.String)
	return data.Bool(ok)
}
