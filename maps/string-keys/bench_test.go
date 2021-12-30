package mapstringkeys

import (
	"fmt"
	"testing"
)

func iterateMapAndReturnValues(m map[string]int, bytesKeys [][]byte) int {
	i := 0
	for _, bytesKey := range bytesKeys {
		key := string(bytesKey)
		i += m[key]
	}
	return i
}

func populateMap(n int) (map[string]int, [][]byte) {
	m := make(map[string]int, n)
	bytesKeys := make([][]byte, 0, len(m))

	for i := 0; i < n; i++ {
		key := fmt.Sprintf("%d", i)
		m[key] = i
		bytesKeys = append(bytesKeys, []byte(key))
	}
	return m, bytesKeys
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
	b.StopTimer()
	m, bytesKeys := populateMap(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sink = iterateMapAndReturnValues(m, bytesKeys)
	}

	if sink == nil {
		b.Fatal("Benchmark did not run")
	}
	sink = (interface{})(nil)
}
