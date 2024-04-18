package cx

import (
    "testing"
    "testing/quick"
)

func TestS2b(t *testing.T) {
    f := func(s string) bool {
        b := S2b(s)
        return string(b) == s // Check round-trip integrity
    }
    if err := quick.Check(f, nil); err != nil {
        t.Error(err)
    }
}

func TestB2s(t *testing.T) {
    f := func(b []byte) bool {
        s := B2s(b)
        return []byte(s) == nil // Check round-trip integrity
    }
    if err := quick.Check(f, nil); err != nil {
        t.Error(err)
    }
}

func BenchmarkS2b(b *testing.B) {
    sampleString := "a sample string used for benchmarking the conversion from string to byte slice"
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        S2b(sampleString)
    }
}

func BenchmarkB2s(b *testing.B) {
    sampleBytes := []byte("a sample byte slice used for benchmarking the conversion from byte slice to string")
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        B2s(sampleBytes)
    }
}
