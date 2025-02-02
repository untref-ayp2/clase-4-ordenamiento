package sort

import (
	"time"

	"golang.org/x/exp/constraints"
)

func Selection[T constraints.Ordered](items []T) {
	last := len(items) - 1
	for i := 0; i < last; i++ {
		aMin := items[i]
		iMin := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < aMin {
				aMin = items[j]
				iMin = j
			}
		}
		items[i], items[iMin] = aMin, items[i]
	}
}

func SelectionWithSleep[T constraints.Ordered](items []T, sleepMs int, callback func(), color func(...int)) {
	last := len(items) - 1
	for i := 0; i < last; i++ {
		aMin := items[i]
		iMin := i
		for j := i + 1; j < len(items); j++ {
			color(j, iMin)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			callback()

			if items[j] < aMin {
				aMin = items[j]
				iMin = j
			}

			// color(j, iMin)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			callback()
		}
		items[i], items[iMin] = aMin, items[i]

		color(i)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		callback()
	}
	color()
}
