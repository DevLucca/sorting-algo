package algorithms

import (
	"time"
)

type BubbleSort struct{}

func (a BubbleSort) Sort(done, changed chan bool, data *[]float64) {
	for sorted := false; !sorted; {
		time.Sleep(time.Second)
		sorted = true
		for i := 0; i < len(*data)-1; i++ {
			time.Sleep(500 * time.Millisecond)
			if (*data)[i] > (*data)[i+1] {
				current, next := (*data)[i], (*data)[i+1]
				(*data)[i], (*data)[i+1] = next, current
				sorted = false
			}
			changed <- true
		}
	}
	done <- true
}
