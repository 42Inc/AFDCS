package distribution

import (
	"math"
)

func Factorial(n int64) int64 {
	var (
		factVal int64 = 1
		i       int64 = 0
	)
	if n < 0 {
		return 0
	} else {
		for i = 1; i <= n; i++ {
			factVal *= int64(i)
		}

	}
	return factVal
}

func Exponential(lambda float64, delta float64) float64 {
	var res float64 = 0.0
	res = 1 - math.Exp(-lambda * delta)
	return res
}

func Erlang(lambda float64, delta float64, n float64) float64 {
	var res float64 = 0.0
	res = lambda * (math.Pow(lambda * delta, n - 1.0)) / float64(Factorial(int64(n - 1.0))) * math.Exp(-lambda * delta)
	return res
}
