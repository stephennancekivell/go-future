package tuple

type T2[A, B any] interface {
	Values() (A, B)
}
type T3[A, B, C any] interface {
	Values() (A, B, C)
}

func New2[A, B any](a A, b B) T2[A, B] {
	return t2[A, B]{a, b}
}

func New3[A, B, C any](a A, b B, c C) T3[A, B, C] {
	return t3[A, B, C]{a, b, c}
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

func (t t2[A, B]) Values() (A, B) {
	return t._1, t._2
}

func (t t3[A, B, C]) Values() (A, B, C) {
	return t._1, t._2, t._3
}
