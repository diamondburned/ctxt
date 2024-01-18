// Package ctxt provides a way to store and retrieve values from a context using
// the type of the value as a key.
package ctxt

import (
	"context"
)

type key[T any] struct{}

// With returns a copy of parent that contains the given value which can be
// retrieved by calling From with the resulting context.
func With[T any](ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, key[T]{}, v)
}

// From returns the value associated with the wanted type.
func From[T any](ctx context.Context) T {
	v, _ := ctx.Value(key[T]{}).(T)
	return v
}

// Has returns true if the context contains a value of the given type.
func Has[T any](ctx context.Context) bool {
	_, ok := ctx.Value(key[T]{}).(T)
	return ok
}
