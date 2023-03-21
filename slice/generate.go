package slice

import (
	"math/rand"
	"time"
)

func GenerateNumbers(size int) []float64 {
	slice := make([]float64, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = float64(rand.Intn(99))
	}
	return slice
}

func GenerateNumbersConsecutive(size int) []float64 {
	slice := make([]float64, size, size)
	for i := 0; i < size; i++ {
		slice[i] = float64(i + 1)
	}
	return slice
}
