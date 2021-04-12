package nqtree

import (
	"errors"
	"math"
)

type Point struct {
	Dimensions []float64
}

func (p1 *Point) Distance(p2 *Point) (float64, error) {
	if len(p1.Dimensions) != len(p2.Dimensions) {
		return 0., errors.New("points dose not have the same number of dimensions")
	}

	var sum float64
	for index, value := range p1.Dimensions {
		sum += math.Pow((value - p2.Dimensions[index]), 2)
	}
	return math.Sqrt(sum), nil
}

type DataPoint struct {
	Data     interface{}
	Position Point
}
