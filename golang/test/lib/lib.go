package lib

func FibonacciNumber(n int64) int64 {
	if n < 2 {
		return n
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}
