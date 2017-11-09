package main

import (
	"testing"
)

func TestMapSearch(t *testing.T) {
	want := [3]string{"scream", "fearful", "cold_sweat"}
	have := MapSearch("frightened")

	if len(have) == 0 {
		t.Error("MapSearch() returned nothing.")
	}

	for i := range have {
		if have[i] != want[i] {
			t.Errorf("Want: %s, have: %s\n", want[i], have[i])
		}
	}
}

func benchmarkMapSearch(s string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapSearch(s)
	}
}

func BenchmarkMapSearch1(b *testing.B) {
	benchmarkMapSearch("happy", b)
}

func BenchmarkMapSearch2(b *testing.B) {
	benchmarkMapSearch("angry", b)
}

func BenchmarkMapSearch3(b *testing.B) {
	benchmarkMapSearch("frightened", b)
}
