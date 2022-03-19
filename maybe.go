package maybe

import "errors"

type An[T any] interface {
	is(T)
}

// Just contains a value.
type Just[T any] struct {
	Value T
}

func (j Just[T]) is(_ T) {}

// Nothing indicates no value is present.
type Nothing[T any] struct{}

func (n Nothing[T]) is(_ T) {}

// Value returns the default value `dflt` if `v` is Nothing. Otherwise it
// returns the result of calling `f` on `v`.
func Value[A, B any](dflt B, f func(a A) B, v An[A]) B {
	switch v.(type) {
	case Just[A]:
		return f(v.(Just[A]).Value)
	case Nothing[A]:
		return dflt
	}

	panic(errors.New("impossible"))
}
