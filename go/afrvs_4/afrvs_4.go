package afrvs_4

// Lecture 8

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"

	"../distribution"
)

var (
	alpha				float64	 = 0
	beta				float64	 = 0
	n           int64   = 0
	k           []int64 = []int64{}
	ModelsCount int64   = 1
	TLimit      float64 = 500
	TimeScale   float64 = 0.0
)

func initFlags() {
	flag.Usage = usage
	flag.Float64Var(&alpha, "alpha", 0.1,
		"ne ebu (alpha)")
	flag.Float64Var(&beta, "beta", 0.5,
		"ne ebu (beta)")
	flag.Int64Var(&n, "n", 5,
		"Reserve size (n)")
	flag.Int64Var(&ModelsCount, "c", 1,
		"Models Count")
	flag.Float64Var(&TLimit, "t", 500,
		"Model time limit")
	flag.Float64Var(&TimeScale, "s", 1,
		"Time Scale")
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [params]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

//
// func MTheor(t float64) float64 {
// 	var (
// 		i   int64   = 0
// 		res float64 = 0.0
// 	)
//
// 	res = alf * (float64(n) + 1) / (2 * (bet - alf))
// 	if res < 0 {
// 		res = 0
// 	}
// 	return res
// }
//
// func DTheor(t float64, M float64) (float64, float64) {
// 	var (
// 		i    int64  = 0
// 		D    float64 = 0.0
// 		C    float64 = 0.0
// 		Up   float64 = 0.0
// 		Down float64 = 0.0
// 	)
// 	// D = (float64(N-n) * lambda * t)
// 	C = math.Pow(float64(i), 2) - float64(i) - 2*float64(N)*(float64(N)-1)*math.Pow(lambda, 2)/((mu+lambda)*(mu+2*lambda)) - 2*(float64(N)-1)*((mu*float64(i)-lambda*(float64(N)-float64(i)))/(mu+lambda))
//
// 	D = 2*float64(N)*(float64(N)-1)*math.Pow(lambda, 2)/((mu+lambda)*(mu+2*lambda)) + C*math.Pow(math.E, -float64((mu+lambda)*t)) + 2*(float64(N)-float64(i))*((mu*float64(i)-lambda*(float64(N)-float64(i)))/(mu+lambda))*math.Pow(math.E, -float64((mu+lambda)*t)) + M - math.Pow(M, 2)
//
// 	if D < 0 {
// 		D = 0
// 	}
// 	Up = M + math.Sqrt(D)
// 	if Up < 0 {
// 		Up = 0
// 	}
// 	Down = M - math.Sqrt(D)
// 	if Down < 0 {
// 		Down = 0
// 	}
// 	return Up, Down
// }

func DPrac(machines []int64, M float64) (float64, float64) {
	var (
		i    int64   = 0
		D    float64 = 0.0
		cnt  int64   = 0
		Up   float64 = 0.0
		Down float64 = 0.0
	)
	for i = 0; i < ModelsCount; i++ {
		cnt = machines[i]
		cnt = int64(math.Pow(float64(cnt), 2))
		D = D + float64(cnt)
	}
	D = (D / float64(ModelsCount)) - math.Pow(M, 2)

	Up = M + math.Sqrt(D)
	if Up < 0 {
		Up = 0
	}
	Down = M - math.Sqrt(D)
	if Down < 0 {
		Down = 0
	}

	return Up, Down
}

func MPrac(machines []int64) float64 {
	var (
		i   int64   = 0
		res float64 = 0.0
		cnt int64   = 0
	)
	for i = 0; i < ModelsCount; i++ {
		cnt = machines[i]
		res = res + float64(cnt)
	}
	res = res / float64(ModelsCount)
	return res
}

func Run() {
	var (
		modelTime float64 = 0.0
		i         int64   = 0
		// MTh       float64 = 0.0
		// DThUp     float64 = 0.0
		// DThDown   float64 = 0.0
		Prob      float64 = 0.0
		MPr       float64 = 0.0
		DPrUp     float64 = 0.0
		DPrDown   float64 = 0.0
	)

	initFlags()

	if ModelsCount < 1 {
		ModelsCount = 1
	}

	log.Printf("Models count: %d\n", ModelsCount)
	FileFP, err := os.OpenFile("data/afrvs_4_FP.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	FileMT, err := os.OpenFile("data/afrvs_4_MT.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	FileDT, err := os.OpenFile("data/afrvs_4_DT.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	FileMP, err := os.OpenFile("data/afrvs_4_MP.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	FileDP, err := os.OpenFile("data/afrvs_4_DP.dat",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	for i = 0; i < ModelsCount; i++ {
		k = append(k, 0)
	}

	for modelTime = 0; modelTime < TLimit; modelTime = modelTime + TimeScale {
		FileFP.WriteString(fmt.Sprintf("%f\t", modelTime))
		for i = 0; i < ModelsCount; i++ {
			if modelTime > 0 {
				FaultProb := distribution.Exponential(alpha, TimeScale)
				RestoreProb := distribution.Erlang(float64(n) * beta, TimeScale, float64(n))
				Prob = rand.Float64()
				if Prob < FaultProb {
					//Fault
					k[i] += n
				}
				Prob = rand.Float64()
				if Prob < RestoreProb {
					//Restore
					k[i]--
					if k[i] < 0 {
						k[i] = 0
					}
				}
			}
			FileFP.WriteString(fmt.Sprintf("%d\t", k[i]))
		}
		FileFP.WriteString("\n")
		MPr = MPrac(k)
		FileMP.WriteString(fmt.Sprintf("%f\t%.6f\n", modelTime, MPr))
		DPrUp, DPrDown = DPrac(k, MPr)
		FileDP.WriteString(fmt.Sprintf("%f\t%.6f\t%.6f\n", modelTime, DPrUp, DPrDown))
	}
	//
	// for i := 0.0; i < modelTime; i = i + TimeScale {
	// 	MTh = MTheor(i)
	// 	DThUp, DThDown = DTheor(i, MTh)
	// 	FileMT.WriteString(fmt.Sprintf("%f\t%.6f\n", i, MTh))
	// 	// MP.WriteString(fmt.Sprintf("%d\t%.6f\n", i, MPrac(dots, i)))
	// 	FileDT.WriteString(fmt.Sprintf("%f\t%.6f\t%.6f\n", i, DThUp, DThDown))
	// }

	FileFP.Close()
	FileMT.Close()
	FileDT.Close()
	FileMP.Close()
	FileDP.Close()
}
