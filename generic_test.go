package main

import (
	"testing"
)

const count = 1_000_000

func BenchmarkNormNdIVec(b *testing.B) {
	for i := 0; i < count; i++ {
		NormNdIVec(1, 2, 3)
	}
}

func BenchmarkNormNd(b *testing.B) {
	for i := 0; i < count; i++ {
		NormNd(1, 2, 3)
	}
}

func BenchmarkNorm2fIVec(b *testing.B) {
	for i := 0; i < count; i++ {
		Norm2fIVec(3, 4)
	}
}

func BenchmarkNorm2f(b *testing.B) {
	for i := 0; i < count; i++ {
		Norm2f(3, 4)
	}
}

func BenchmarkNorm2fSimple(b *testing.B) {
	for i := 0; i < count; i++ {
		Norm2fSimple(3, 4)
	}
}

func BenchmarkNorm2fAny(b *testing.B) {
	for i := 0; i < count; i++ {
		Norm2fAny(3, 4)
	}
}
