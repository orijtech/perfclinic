package fmtcastbefore

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func formatIt(b []byte) string {
	return fmt.Sprintf("value: %s", b)
}

var values = []struct {
	in   []byte
	want string
}{
	{
		in:   []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		want: "value: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	},
	{
		in:   []byte(""),
		want: "value: ",
	},
	{
		in:   []byte("abc"),
		want: "value: abc",
	},
	{
		in:   bytes.Repeat([]byte("🍾😎"), 1e2),
		want: "value: " + strings.Repeat("🍾😎", 1e2),
	},
}

var sink interface{}

func BenchmarkFormatIt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nBytes := 0
		for _, value := range values {
			if got := formatIt(value.in); got != value.want {
				b.Fatalf("Result mismatch\n\tgot:  %q\n\twant: %q", got, value.want)
			}
			nBytes += len(value.in)
		}
		b.SetBytes(int64(nBytes))
                sink = nBytes
	}
	if sink == nil {
		b.Fatal("Benchmark didn't run")
	}
	// Reset the sink.
	sink = (interface{})(nil)
}
