package sillymath

func Add(a, b int) int {
	return a + b
}

func Square(a int) int {
	return a * a
}

// Fib(n) = Fib(n-2) + Fib(n-1)
// Fib(0) = 0
// Fib(1) = 1
//
// Sequence like: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89
func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return Fib(n-2) + Fib(n-1)
}

func FibFaster(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	fibs := []int{0, 1}
	for i := 2; i <= n; i++ {
		fibs = append(fibs, fibs[i-2]+fibs[i-1])
	}
	return fibs[n]
}

func FibEvenFaster(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	var theFib int
	fib0 := 0
	fib1 := 1
	for i := 2; i <= n; i++ {
		theFib = fib0 + fib1
		fib0 = fib1
		fib1 = theFib
	}
	return theFib
}
