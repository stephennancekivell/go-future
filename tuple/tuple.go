package tuple

// A Tuple is a few values put together.
// create a tuple with `tuple.New2(value1, value2)`
// then get the values out with `.Values()`

type T2[A, B any] interface {
	Values() (A, B)
}
type T3[A, B, C any] interface {
	Values() (A, B, C)
}
type T4[A, B, C, D any] interface {
	Values() (A, B, C, D)
}

func New2[A, B any](a A, b B) T2[A, B] {
	return t2[A, B]{a, b}
}

func New3[A, B, C any](a A, b B, c C) T3[A, B, C] {
	return t3[A, B, C]{a, b, c}
}
func New4[A, B, C, D any](a A, b B, c C, d D) T4[A, B, C, D] {
	return t4[A, B, C, D]{a, b, c, d}
}

type t2[A, B any] struct {
	_1 A
	_2 B
}
type t3[A, B, C any] struct {
	_1 A
	_2 B
	_3 C
}

type t4[A, B, C, D any] struct {
	_1 A
	_2 B
	_3 C
	_4 D
}

func (t t2[A, B]) Values() (A, B) {
	return t._1, t._2
}

func (t t3[A, B, C]) Values() (A, B, C) {
	return t._1, t._2, t._3
}

func (t t4[A, B, C, D]) Values() (A, B, C, D) {
	return t._1, t._2, t._3, t._4
}
