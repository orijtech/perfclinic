package sliceclearing

import "testing"

func returnValuesAndClearSlice(s []int) int {
	tot := 0
	l := len(s)
	for i := 0; i < l; i++ {
		s[i] = 0
	}
	return tot
}

func populateSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

var sink interface{}

func benchmarkIt(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		s := populateSlice(n)
		sink = returnValuesAndClearSlice(s)
	}

	if sink == nil {
		b.Fatal("Benchmark did not run")
	}
	sink = (interface{})(nil)
}

func BenchmarkIt10(b *testing.B) {
	benchmarkIt(b, 10)
}

func BenchmarkIt100(b *testing.B) {
	benchmarkIt(b, 100)
}

func BenchmarkIt1000(b *testing.B) {
	benchmarkIt(b, 1000)
}

func BenchmarkIt10000(b *testing.B) {
	benchmarkIt(b, 10000)
}

func BenchmarkIt100000(b *testing.B) {
	benchmarkIt(b, 100000)
}

func BenchmarkIt1000000(b *testing.B) {
	benchmarkIt(b, 1000000)
}
