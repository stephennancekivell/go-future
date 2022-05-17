package future

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stephennancekivell/go-future/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNew_makes_new_literal(t *testing.T) {
	assert := assert.New(t)

	futureValue := New(func() string {
		return "value"
	})

	assert.Equal(futureValue.Get(), "value", "gat get it")
	assert.Equal(futureValue.Get(), "value", "can do it again be same")
}

func TestNew_makes_new_struct(t *testing.T) {
	assert := assert.New(t)

	type t_value struct{ value string }

	futureValue := New(func() t_value {
		return t_value{"value"}
	})

	assert.Equal(futureValue.Get(), t_value{"value"}, "gat get it")
	assert.Equal(futureValue.Get(), t_value{"value"}, "can do it again be same")
}

func TestNew_makes_new_struct_ptr(t *testing.T) {
	assert := assert.New(t)

	type t_value struct{ value string }

	futureValue := New(func() *t_value {
		return &t_value{"value"}
	})

	assert.Equal(*futureValue.Get(), t_value{"value"}, "gat get it")
	assert.Equal(*futureValue.Get(), t_value{"value"}, "can do it again be same")
	assert.Equal(futureValue.Get(), futureValue.Get(), "pointers are the same")
}

func TestPure(t *testing.T) {
	assert := assert.New(t)

	f := Pure("value")

	assert.Equal(f.Get(), "value", "should be equal")
	assert.Equal(f.Get(), "value", "can get twice")
}

func TestSequence(t *testing.T) {
	assert := assert.New(t)

	f := Sequence(
		[]Future[string]{
			Pure("one"),
			Pure("two"),
			Pure("three"),
		},
	)

	assert.Equal(f.Get(), []string{"one", "two", "three"}, "should be equal")
}

func TestTraverse(t *testing.T) {
	assert := assert.New(t)

	f := Traverse(
		[]string{
			"one",
			"two",
			"three",
		},
		func(value string) Future[string] {
			return Pure(value)
		},
	)

	assert.Equal(f.Get(), []string{"one", "two", "three"}, "should be equal")

}

func TestRace_one(t *testing.T) {
	assert := assert.New(t)

	f := Race(
		New(func() string { return "one" }),
	)

	assert.Equal(f.Get(), "one", "should be the value")
}

func TestRace_two_first_slow(t *testing.T) {
	assert := assert.New(t)

	var mu sync.Mutex
	mu.Lock()

	f := Race(
		New(func() string { mu.Lock(); return "one" }),
		New(func() string { return "two" }),
	)

	assert.Equal(f.Get(), "two", "should be the value")
	mu.Unlock()

}

func TestFutureTuple(t *testing.T) {
	assert := assert.New(t)

	futureValue := New(func() tuple.T2[string, error] {
		return tuple.New2("value", fmt.Errorf("wow"))
	})

	v, e := futureValue.Get().Values()
	assert.Equal(v, "value")
	assert.Equal(e, fmt.Errorf("wow"))
}
