package lib

import (
	"testing"
)

func TestFibonacciNumber1(t *testing.T) {
	r := FibonacciNumber(10)
	if r != 55 {
		t.Errorf("FibonacciNumber(10) failed. Got %d, expected 3.", r)
	}
}
