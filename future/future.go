package future

import (
	"sync"
)

// Future represents a value that might not be ready yet.
//
// Create futures with New(func() ... )
//
// Get the value out with f.Get()
type Future[T any] interface {
	Get() T
}

type futureImpl[T any] struct {
	wg    sync.WaitGroup
	value T
}

func (f *futureImpl[T]) Get() T {
	f.wg.Wait()

	return f.value
}

// Creates a new Future running fn in a go routine.
func New[T any](fn func() T) Future[T] {

	f := futureImpl[T]{}
	f.wg.Add(1)

	go func() {
		f.value = fn()
		f.wg.Done()
	}()

	return &f
}

type pureFutureImpl[T any] struct {
	value T
}

// Creates a new Future from the values.
func Pure[T any](value T) Future[T] {
	return pureFutureImpl[T]{value}
}
func (f pureFutureImpl[T]) Get() T {
	return f.value
}

func Sequence[T any](xs []Future[T]) Future[[]T] {

	return New(
		func() []T {
			ret := make([]T, len(xs))

			for i := 0; i < len(xs); i++ {
				ret[i] = xs[i].Get()
			}
			return ret
		},
	)
}

// TODO TraversePool, TraverseBonded
func Traverse[A any, B any](xs []A, fn func(value A) Future[B]) Future[[]B] {

	futures := make([]Future[B], len(xs))

	for i, a := range xs {
		futures[i] = fn(a)
	}

	return Sequence(futures)
}

func SequenceFlat[T any](xs []Future[[]T]) Future[[]T] {
	return New(
		func() []T {
			f := Sequence(xs)
			xxs := f.Get()

			ret := []T{}

			for _, xs := range xxs {
				ret = append(ret, xs...)
			}
			return ret
		},
	)
}

func Race[T any](head Future[T], rest ...Future[T]) Future[T] {
	return New(func() T {

		ch := make(chan T, len(rest)+1)

		go func() {
			ch <- head.Get()
		}()

		for i := 0; i < len(rest); i++ {
			i_ := i
			go func() {
				ch <- rest[i_].Get()
			}()

		}

		return <-ch
	})
}
