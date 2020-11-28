package main

import (
	"math"
	"testing"
)

var g float32

func BenchmarkSqrt(b *testing.B) {
	b.Run("baseline", func(b *testing.B) {
		g = 0
		for i := 0; i < b.N; i++ {
			v := float64(i)
			g += float32(v)
		}
	})

	b.Run("built-in", func(b *testing.B) {
		g = 0
		for i := 0; i < b.N; i++ {
			v := float64(i)
			g += float32(math.Sqrt(v + 100.137*v))
		}
	})

	b.Run("avx256", func(b *testing.B) {
		g = 0
		var a1 [batchSize]float32
		var a2 [batchSize]float32
		for i := 0; i < b.N; i += batchSize {
			for j := 0; j < batchSize; j++ {
				v := float32(i + j)
				a1[j] = v + 100.137*v
			}
			SqrtAvx256(a1[:], a2[:])
			for j := 0; j < batchSize; j++ {
				g += a2[j]
			}
		}
	})

	b.Run("avx256-prepared", func(b *testing.B) {
		g = 0
		const sizePrepared = batchSize * 1024 * 64
		var a1 [sizePrepared]float32
		var a2 [batchSize]float32
		for i := 0; i < len(a1); i++ {
			v := float32(i)
			a1[i] = v + 100.137*v
		}

		b.ResetTimer()
		for i := 0; i < b.N; i += batchSize {
			SqrtAvx256(a1[i&(sizePrepared-1):], a2[:])
			for j := 0; j < batchSize; j++ {
				g += a2[j]
			}
		}
	})
}
