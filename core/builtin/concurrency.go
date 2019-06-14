package builtin

import (
	"gitlab.com/kode4food/ale/data"
	"gitlab.com/kode4food/ale/stdlib"
)

const (
	// ChannelType is the type name for a channel
	ChannelType = data.String("channel")

	// EmitKey is the key used to emit to a Channel
	EmitKey = data.Keyword("emit")

	// SequenceKey is the key used to retrieve the Sequence from a Channel
	SequenceKey = data.Keyword("seq")
)

// Go runs the provided function asynchronously
func Go(args ...data.Value) data.Value {
	fn := args[0].(data.Caller)
	restArgs := args[1:]
	go fn.Caller()(restArgs...)
	return data.Null
}

// Chan instantiates a new go channel
func Chan(args ...data.Value) data.Value {
	var size int
	if len(args) != 0 {
		size = int(args[0].(data.Integer))
	}
	e, s := stdlib.NewChannel(size)

	return data.Object{
		data.TypeKey: ChannelType,
		EmitKey:      bindWriter(e),
		CloseKey:     bindCloser(e),
		SequenceKey:  s,
	}
}

// Promise instantiates a new eventually-fulfilled promise
func Promise(args ...data.Value) data.Value {
	if len(args) == 0 {
		return stdlib.NewPromise()
	}
	p := stdlib.NewPromise()
	p.Deliver(args[0])
	return p
}

// IsPromise returns whether or not the specified value is a promise
func IsPromise(args ...data.Value) data.Value {
	_, ok := args[0].(stdlib.Promise)
	return data.Bool(ok)
}

// IsDelivered returns whether or not the specified promise has been delivered
func IsDelivered(args ...data.Value) data.Value {
	p := args[0].(stdlib.Promise)
	return data.Bool(p.IsDelivered())
}
