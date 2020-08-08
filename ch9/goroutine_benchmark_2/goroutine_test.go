package main

import (
	"testing"
	"time"
)

func TestPrint2(t *testing.T) {
	print2()
}

func TestGoPrint2(t *testing.T) {
	goPrint2()
	time.Sleep(1 * time.Millisecond)
}

func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}
