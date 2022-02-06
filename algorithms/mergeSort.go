package algorithms

import (
	"fmt"
	"time"
)

type MergeSort struct{}

func (m MergeSort) Sort(done, changed chan bool, data *[]float64) {
	fmt.Println("sort", &data)
	time.Sleep(5 * time.Second)
	*data = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("merge", &data)
	changed <- true
	time.Sleep(15 * time.Second)

	*data = []float64{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	fmt.Println("merge", &data)
	changed <- true
	time.Sleep(15 * time.Second)
	done <- true
}
