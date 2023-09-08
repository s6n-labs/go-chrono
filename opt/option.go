package opt

import (
	"github.com/s6n-labs/go-chrono/tup"
)

type Option[T any] struct {
	value  T
	isSome bool
}

func (o Option[T]) IsSome() bool {
	return o.isSome
}

func (o Option[T]) IsNone() bool {
	return !o.isSome
}

func (o Option[T]) Unwrap() T {
	return o.value
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value:  value,
		isSome: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func AndThen[T, U any](self Option[T], fn func(T) Option[U]) Option[U] {
	if self.IsNone() {
		return None[U]()
	}

	return fn(self.Unwrap())
}

func Map[T, U any](self Option[T], fn func(T) U) Option[U] {
	return AndThen[T, U](self, func(t T) Option[U] {
		return Some(fn(t))
	})
}

func Zip[T, U any](self Option[T], rhs Option[U]) Option[tup.Tuple[T, U]] {
	if self.IsNone() || rhs.IsNone() {
		return None[tup.Tuple[T, U]]()
	}

	return Some(tup.NewTuple(self.Unwrap(), rhs.Unwrap()))
}
