package randArray

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateArray(lenght int) []float64 {
	array := make([]float64, lenght)
	for i := int(0); i < lenght; i++ {
		array[i] = float64(rand.Intn(lenght * 10))
	}
	return array
}
