package lib

import (
	"testing"
)

//功能测试
func TestFibonacci(t *testing.T) {
	r := Fibonacci(10)
	if r != 55 {
		t.Errorf("Fibonacci(10) failed. Got %d, expected 55.", r)
	}
}

//性能测试
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}
