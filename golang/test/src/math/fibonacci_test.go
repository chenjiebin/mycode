package math

import (
	"testing"
)

func TestFibonacciNumber1(t *testing.T) {
	r := FibonacciNumber(10)
	if r != 54 {
		t.Errorf("Fibonacci(10) failed. Got %d, expected 55.", r)
	}
}
