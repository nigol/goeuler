package problem2

func fib() int {
	sum := 0
	a := 1
	b := 2
	for b < 4000000 {
		if b%2 == 0 {
			sum = sum + b
		}
		c := a + b
		a = b
		b = c
	}
	return sum
}
