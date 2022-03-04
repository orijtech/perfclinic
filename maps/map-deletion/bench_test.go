package mapclearing

import "testing"

func returnValuesAndClearMap(m map[int]int) int {
	i := 0
	// 1. Just get the values.
	for _, value := range m {
		i += value
	}

	// 2. Use the map clearing idiom solely now.
	for key := range m {
		delete(m, key)
	}
	return i
}

func populateMap(n int) map[int]int {
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		m[i] = i
	}
	return m
}

var sink interface{}

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

func benchmarkIt(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := populateMap(n)
		sink = returnValuesAndClearMap(m)
	}

	if sink == nil {
		b.Fatal("Benchmark did not run")
	}
	sink = (interface{})(nil)
}
