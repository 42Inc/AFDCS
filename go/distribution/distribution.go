package distribution

import (
	"math"
	"math/rand"
)

func Exponential(lambda float64) float64 {
	var res float64 = 0.0
	res = -math.Log(1-rand.Float64()) / lambda
	return res
}
