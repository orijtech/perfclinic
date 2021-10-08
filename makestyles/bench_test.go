package makestyles

import "testing"

func makeIt(n int) (data []int) {
	data = make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}

var sink interface{}

func benchmarkMake(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		sink = makeIt(size)
	}
	if sink == nil {
		b.Fatal("Benchmark didn't run")
	}
}

func BenchmarkMake10(b *testing.B) {
	benchmarkMake(b, 10)
}

func BenchmarkMake100(b *testing.B) {
	benchmarkMake(b, 100)
}

func BenchmarkMake1000(b *testing.B) {
	benchmarkMake(b, 1000)
}

func BenchmarkMake10000(b *testing.B) {
	benchmarkMake(b, 10000)
}

func BenchmarkMake100000(b *testing.B) {
	benchmarkMake(b, 100000)
}

func BenchmarkMake1000000(b *testing.B) {
	benchmarkMake(b, 1000000)
}
