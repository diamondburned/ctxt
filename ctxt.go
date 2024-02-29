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
func From[T any](ctx context.Context) (T, bool) {
	v, ok := ctx.Value(key[T]{}).(T)
	return v, ok
}

// FromOr returns the value associated with the wanted type or the given default
// value if the type is not found.
func FromOr[T any](ctx context.Context, def T) T {
	v, ok := From[T](ctx)
	if !ok {
		return def
	}
	return v
}

// FromOrFunc returns the value associated with the wanted type or the result of
// the given function if the type is not found.
// For example:
//
//	logger := ctxt.FromOrFunc(ctx, slog.Default)
func FromOrFunc[T any](ctx context.Context, f func() T) T {
	v, ok := From[T](ctx)
	if !ok {
		return f()
	}
	return v
}
