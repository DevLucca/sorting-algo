package algorithms

import "time"

type SelectionSort struct{}

func (s SelectionSort) Sort(done, changed chan bool, data *[]float64) {
	minValue := 0.0
	minValueIndex := 0
	for i := 0; i < len(*data); i++ {
		time.Sleep(time.Second)
		currentValue := (*data)[i]
		minValue = currentValue
		for j := i; j < len(*data); j++ {
			internalCurrentValue := (*data)[j]
			if minValue >= internalCurrentValue {
				minValue = internalCurrentValue
				minValueIndex = j
			}
		}
		(*data)[i] = minValue
		(*data)[minValueIndex] = currentValue
		changed <- true
	}

	done <- true
}
