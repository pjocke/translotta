package main

import (
	"testing"
)

func TestStructSearch(t *testing.T) {
	want := [3]string{"scream", "fearful", "cold_sweat"}
	have := StructSearch("frightened")

	if len(have) == 0 {
		t.Error("StructSearch() returned nothing.")
	}

	for i := range have {
		if have[i] != want[i] {
			t.Errorf("Want: %s, have: %s\n", want[i], have[i])
		}
	}
}

func benchmarkStructSearch(s string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		StructSearch(s)
	}
}

func BenchmarkStructSearch1(b *testing.B) {
	benchmarkStructSearch("happy", b)
}

func BenchmarkStructSearch2(b *testing.B) {
	benchmarkStructSearch("angry", b)
}

func BenchmarkStructSearch3(b *testing.B) {
	benchmarkStructSearch("frightened", b)
}
