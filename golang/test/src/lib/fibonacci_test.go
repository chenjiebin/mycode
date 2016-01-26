package lib

import (
	"fmt"
	"testing"
)

//功能测试
//func TestFibonacci(t *testing.T) {
//	r := Fibonacci(10)
//	if r != 55 {
//		t.Errorf("Fibonacci(10) failed. Got %d, expected 55.", r)
//	}
//}

////性能测试
//func BenchmarkFibonacci(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Fibonacci(10)
//	}
//}

func ExampleFibonacci() {
	fmt.Println(Fibonacci(10))
	// Output: 55
}

//测试参数为5的性能
func BenchmarkFibonacci5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(5)
	}
}

//测试参数为20的性能
func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}
