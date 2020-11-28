package main

import (
	"fmt"
	"math"
)

const batchSize = 256 / 32

func main() {
	var c int
	AddTwoInts(100, 200, &c)
	fmt.Println("c:", c)

	x := make([]float32, batchSize)
	for i := 0; i < batchSize; i++ {
		x[i] = float32(i)
	}

	r := make([]float32, batchSize)
	SqrtAvx256(x, r)
	for i := 0; i < batchSize; i++ {
		fmt.Printf("i: %d, x: %f, r: %f, math.Sqrt: %f\n",
			i, x[i], r[i], math.Sqrt(float64(x[i])))
	}
}
