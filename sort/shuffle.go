package sort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func Shuffle[T constraints.Ordered](items []T) {
	for i := range items {
		j := rand.Intn(i + 1)
		items[i], items[j] = items[j], items[i]
	}
}
