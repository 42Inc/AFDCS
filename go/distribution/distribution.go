package distribution

import (
	"math"
)

func Exponential(lambda float64, delta float64) float64 {
	var res float64 = 0.0
	res = 1 - math.Exp(-lambda * delta)
	return res
}
