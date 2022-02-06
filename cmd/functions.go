package cmd

import (
	"github.com/DevLucca/sorting-algo/algorithms"
)

var mapSupportedAlgorithms = map[string]Algorithm{
	"bubble":    algorithms.BubbleSort{},
	"insertion": algorithms.InsertionSort{},
	"selection": algorithms.SelectionSort{},
	"quick":     algorithms.QuickSort{},
	"merge":     algorithms.MergeSort{},
	"heap":      algorithms.HeapSort{},
}

type context struct {
	Strategy Algorithm
}

type Algorithm interface {
	Sort(chan bool, chan bool, chan []float64)
}

func (c *context) SetStrategy(s string) {
	c.Strategy = mapSupportedAlgorithms[s]
}

func (c *context) ExecuteStrategy(done, changed chan bool, data chan []float64) (err error) {
	if c.Strategy != nil {
		c.Strategy.Sort(done, changed, data)
	}
	return
}
