package sliceclearing

import "testing"

type foo struct {
	data []int
}

func returnValuesAndClearSlice(f *foo) int {
	tot := 0
	for _, val := range f.data {
		tot += val
	}
	f.data = make([]int, len(f.data))
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
		f := &foo{}
		f.data = populateSlice(n)
		sink = returnValuesAndClearSlice(f)
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
