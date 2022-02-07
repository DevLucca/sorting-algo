package algorithms

import (
	"time"
)

type InsertionSort struct{}

func (i InsertionSort) Sort(done, changed chan bool, data *[]float64) {
	for idx, v := range *data {
		time.Sleep(time.Second)
		for idx > 0 && v < (*data)[idx-1] {
			(*data)[idx] = (*data)[idx-1]
			idx--
		}
		(*data)[idx] = v
		changed <- true
	}

	done <- true
}
