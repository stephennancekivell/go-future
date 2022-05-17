package future

import (
	"testing"
)

func get100Test() {
	f := New(func() big { return big{} })

	for i := 0; i < 1000; i++ {
		f.Get()
	}
}

type big struct {
	a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p string
}

func BenchmarkNewGet100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		get100Test()
	}
}
