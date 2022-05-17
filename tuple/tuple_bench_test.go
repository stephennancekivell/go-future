package tuple

import "testing"

func noop(v string, e error) {}

func multiRet() {
	f := func() (string, error) {
		return "v", nil
	}

	for i := 0; i < 1000; i++ {
		noop(f())
	}
}

func BenchmarkMultiReturn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		multiRet()
	}
}

func tupleRet() {
	f := func() T2[string, error] {
		return New2[string, error]("v", nil)
	}

	for i := 0; i < 1000; i++ {
		v := f()
		noop(v.Values())
	}
}

func BenchmarkTuple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tupleRet()
	}
}
