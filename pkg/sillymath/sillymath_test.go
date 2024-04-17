package sillymath

import (
	"fmt"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	a, b := 1, 2

	got := Add(a, b)
	want := a + b

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFib(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	for i, n := range fibs {
		t.Run(fmt.Sprintf("Fib(%d)", i), func(t *testing.T) {
			t.Parallel()

			got := Fib(i)
			want := n

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

func TestFibFaster(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	for i, n := range fibs {
		t.Run(fmt.Sprintf("FibFaster(%d)", i), func(t *testing.T) {
			t.Parallel()

			got := FibFaster(i)
			want := n

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

func TestFibEvenFaster(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	for i, n := range fibs {
		t.Run(fmt.Sprintf("FibEvenFaster(%d)", i), func(t *testing.T) {
			t.Parallel()

			got := FibEvenFaster(i)
			want := n

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

const (
	START = 15
	END   = 20
)

func BenchmarkFib(b *testing.B) {
	for n := START; n < END; n++ {
		b.Run(fmt.Sprintf("Fib(%d)", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fib(n)
			}
		})
	}
}

func BenchmarkFibFaster(b *testing.B) {
	for n := START; n < END; n++ {
		b.Run(fmt.Sprintf("FibFaster(%d)", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibFaster(n)
			}
		})
	}
}

func BenchmarkFibEvenFaster(b *testing.B) {
	for n := START; n < END; n++ {
		b.Run(fmt.Sprintf("FibEvenFaster(%d)", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibEvenFaster(n)
			}
		})
	}
}

func TestSquare(t *testing.T) {
	a := 4

	got := Square(a)
	want := a * a

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func FuzzSquare(f *testing.F) {
	// seed corpus
	vals := []int{1, 2, 3, 4, 5, 6}
	for _, v := range vals {
		f.Add(v)
	}

	// run fuzz test
	f.Fuzz(func(t *testing.T, v int) {
		squared := Square(v)
		root := int(math.Sqrt(float64(squared)))

		got := root
		want := v
		if v < 0 {
			want = -want
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
