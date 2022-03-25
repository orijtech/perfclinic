package maps

import (
	"testing"
)

var sink interface{}

func BenchmarkCopyIt10(b *testing.B) {
	benchmarkCopyIt(b, 10)
}

func BenchmarkCopyIt100(b *testing.B) {
	benchmarkCopyIt(b, 100)
}

func BenchmarkCopyIt1000(b *testing.B) {
	benchmarkCopyIt(b, 1000)
}

func BenchmarkCopyIt10000(b *testing.B) {
	benchmarkCopyIt(b, 10000)
}

func BenchmarkCopyIt100000(b *testing.B) {
	benchmarkCopyIt(b, 100000)
}

func benchmarkCopyIt(b *testing.B, n int) {
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = i + 1
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sink = copyIt(m)
	}
	if sink == nil {
		b.Fatal("Benchmark didn't run")
	}
	sink = (interface{})(nil)
}
