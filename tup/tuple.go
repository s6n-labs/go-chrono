package tup

type Tuple[T0, T1 any] struct {
	V0 T0
	V1 T1
}

func (t Tuple[T0, T1]) Explode() (T0, T1) {
	return t.V0, t.V1
}

func NewTuple[T0, T1 any](v0 T0, v1 T1) Tuple[T0, T1] {
	return Tuple[T0, T1]{
		V0: v0,
		V1: v1,
	}
}

type Tuple3[T0, T1, T2 any] struct {
	V0 T0
	V1 T1
	V2 T2
}

func (t Tuple3[T0, T1, T2]) Explode() (T0, T1, T2) {
	return t.V0, t.V1, t.V2
}

func NewTuple3[T0, T1, T2 any](v0 T0, v1 T1, v2 T2) Tuple3[T0, T1, T2] {
	return Tuple3[T0, T1, T2]{
		V0: v0,
		V1: v1,
		V2: v2,
	}
}
