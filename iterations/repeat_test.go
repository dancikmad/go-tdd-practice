package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeating abc, 3 times", func(t *testing.T) {
		repeated := ExampleReapeat("abc", 3)
		expected := "abcabcabc"

		if repeated != expected {
			t.Errorf("expected %q but not %q", expected, repeated)
		}
	})
	t.Run("repeating a, 5 times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but not %q", expected, repeated)
		}
	})
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
		ExampleReapeat("abc", 3)
	}
}
