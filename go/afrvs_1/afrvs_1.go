package afrvs_1

import (
	"fmt"
	"log"
	"math"
	"os"

	"../distribution"
)

var (
	mu     float64 = 0
	n      int64   = 20
	N      int64   = 1E+3 + n
	lambda float64 = 1E-4
	k      int64   = n
)

func faultPolicy(lambda float64) float64 {
	// return 1.0 / (float64(N-n) * lambda)
	return distribution.Exponential(lambda)
}

func MTheor(t int64) float64 {
	var res float64 = (float64(n) - float64(N-n)*lambda*float64(t))
	if res < 0 {
		res = 0
	}
	return res
}

func MPrac(dots [][]int64, t int64) float64 {
	var (
		res float64 = 0.0
		_t  int64   = t
		min int64   = 0
		max int64   = 0
	)
	if len(dots) > 0 {
		for i := range dots {
			if i > 0 {
				if (dots[i-1][0] < _t) && (_t <= dots[i][0]) {
					min = int64(i - 1)
					max = int64(i)
				}
			} else {
				if _t < dots[i][0] {
					min = 0
					max = 0
				}
			}
		}
		if min > 0 && max > 0 {
			for i := 0; int64(i) < max; i++ {
				if min > int64(i) {
					if i > 0 {
						res = res + float64((dots[i][1]+1)*(dots[i][0]-dots[i-1][0]))
					} else {
						res = res + float64((dots[i][1]+1)*(dots[i][0]))
					}
				} else {
					if i > 0 {
						res = res + float64((dots[i][1]+1)*(_t-dots[i-1][0]))
					} else {
						res = res + float64((dots[i][1]+1)*(_t))
					}
				}
			}
		} else {
			res = res + float64((dots[0][1]+1)*(_t))
		}
		res = res / float64(t)
	} else {
		res = float64(_t)
	}

	if t == 0 {
		res = float64(n)
	}

	return res
}

func DTheor(t int64) float64 {
	var res float64 = (float64(N-n) * lambda * float64(t))
	if res < 0 {
		res = 0
	}
	return res
}

func Run() {
	var (
		timeToFault   int64     = int64(faultPolicy(lambda))
		modelTime     int64     = 0
		lastFaultTime int64     = 0
		dots          [][]int64 = [][]int64{}
		i             int64     = 0
		MTh           float64   = 0.0
		DTh           float64   = 0.0
		TLimit        int64     = 5000
		limit         bool      = false
	)

	FP, err := os.OpenFile("data/afrvs_1_FP.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	MT, err := os.OpenFile("data/afrvs_1_MT.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	DT, err := os.OpenFile("data/afrvs_1_DT.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	MP, err := os.OpenFile("data/afrvs_1_MP.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Time\tleft\tnFault\n")
	for (modelTime < TLimit && limit) || (k > 0 && !limit) {
		if modelTime-lastFaultTime >= timeToFault && k > 0 {
			lastFaultTime = modelTime
			timeToFault = int64(faultPolicy(lambda))
			k--
			fmt.Printf("%d\t%d\t%d\n", modelTime, k, timeToFault)
			FP.WriteString(fmt.Sprintf("%d\t%d\n", modelTime, k))
			dots = append(dots, []int64{modelTime, k})
		}
		modelTime++

	}

	for i := range dots {
		fmt.Printf("%d\t%d\n", dots[i][0], dots[i][1])
	}

	for i = 0; i < modelTime; i++ {
		MTh = MTheor(i)
		DTh = DTheor(i)
		MT.WriteString(fmt.Sprintf("%d\t%.6f\n", i, MTh))
		MP.WriteString(fmt.Sprintf("%d\t%.6f\n", i, MPrac(dots, i)))
		DT.WriteString(fmt.Sprintf("%d\t%.6f\t%.6f\n", i, MTh+math.Sqrt(DTh), MTh-math.Sqrt(DTh)))
	}

	FP.Close()
	MT.Close()
	DT.Close()
	MP.Close()
}
