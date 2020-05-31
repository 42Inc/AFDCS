package afrvs_1

import (
	"fmt"
	"log"
	"os"

	"../distribution"
)

var (
	lambda float64 = 1E-4
	mu     float64 = 0
	n      int64   = 20
	N      int64   = 1E+3 + n
	k      int64   = n
)

func MTheor(t int64) float64 {
	var res float64 = (float64(n) - float64((N-n))*lambda*float64(t))
	if res < 0 {
		res = 0
	}
	return res
}

func MPrac(dots [][]int64, t int64) float64 {
	var (
		res float64 = 0.0
		i int64 = 0
		dot int64 = 0
	)
	for i = 0; i < t; i++ {
		if (i >= dots[dot][0]) {
			dot++
		}
		res = res + float64(dots[dot][1] + 1)
	}
	res = res / float64(t)
	if (t == 0) {
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
		timeToFault   int64     = int64(distribution.Exponential(lambda))
		modelTime     int64     = 0
		lastFaultTime int64     = 0
		dots          [][]int64 = [][]int64{}
		M             float64   = float64(k)
		D             float64   = 0
		i             int64     = 0
	)

	FP, err := os.OpenFile("afrvs_1_FP.dat", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666)
	if err != nil {
		log.Fatal(err)
	}
	MT, err := os.OpenFile("afrvs_1_MT.dat", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666)
	if err != nil {
		log.Fatal(err)
	}
	DT, err := os.OpenFile("afrvs_1_DT.dat", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666)
	if err != nil {
		log.Fatal(err)
	}
	MP, err := os.OpenFile("afrvs_1_MP.dat", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Time\tleft\tnFault\tM\tD\t\n")
	for k > 0 {
		if modelTime-lastFaultTime >= timeToFault {
			lastFaultTime = modelTime
			timeToFault = int64(distribution.Exponential(lambda))
			k--
			fmt.Printf("%d\t%d\t%d\t%f\t%f\t\n", modelTime, k, timeToFault, M, D)
			FP.WriteString(fmt.Sprintf("%d\t%d\n", modelTime, k))
			dots = append(dots, []int64{modelTime, k})
		}

		modelTime++
	}

	for i := range dots {
		fmt.Printf("%d\t%d\n", dots[i][0], dots[i][1])
	}

	for i = 0; i < modelTime; i++ {
		MT.WriteString(fmt.Sprintf("%d\t%.6f\n", i, MTheor(i)))
		MP.WriteString(fmt.Sprintf("%d\t%.6f\n", i, MPrac(dots, i)))
		DT.WriteString(fmt.Sprintf("%d\t%.6f\n", i, DTheor(i)))
	}


	FP.Close()
	MT.Close()
	DT.Close()
	MP.Close()
}
