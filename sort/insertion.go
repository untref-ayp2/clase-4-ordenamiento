package sort

import (
	"golang.org/x/exp/constraints"
	"time"
)

func Insertion[T constraints.Ordered](items []T) {
	for i := 1; i < len(items); i++ {
		value := items[i]
		j := i - 1
		for j >= 0 && items[j] > value {
			items[j+1] = items[j]
			j = j - 1
		}
		items[j+1] = value
	}
}

func InsertionWithSleep[T constraints.Ordered](items []T, sleepMs int, callback func(), color func(...int)) {
	for i := 1; i < len(items); i++ {
		value := items[i]
		j := i - 1
		for j >= 0 && items[j] > value {
			color(j, j+1)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			callback()

			items[j+1] = items[j]
			j = j - 1

			color(j+1, j+2)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			callback()
		}
		items[j+1] = value

		color(j + 1)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		callback()
	}
	color()
}
