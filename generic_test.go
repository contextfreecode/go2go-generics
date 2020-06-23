package main

import (
	"testing"
)

func BenchmarkNormNdIVec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormNdIVec(1, 2, 3)
	}
}

func BenchmarkNormNd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormNd(1, 2, 3)
	}
}

func BenchmarkNorm2fIVec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Norm2fIVec(3, 4)
	}
}

func BenchmarkNorm2f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Norm2f(3, 4)
	}
}

func BenchmarkNorm2fSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Norm2fSimple(3, 4)
	}
}

func BenchmarkNorm2fAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Norm2fAny(3, 4)
	}
}
