package testing_test

import "testing"

func Func1() {
}

func Func2() {
}

func BenchmarkFunc1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Func1()
	}
}

func BenchmarkFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Func2()
	}
}
