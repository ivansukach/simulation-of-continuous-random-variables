package distributions

import "math"

func LogisticDistributionVariates(m float64, k float64, basicVariate []float64) []float64 {
	n := len(basicVariate)
	logisticVariates := make([]float64, n)
	for i := 0; i < n; i++ {
		logisticVariates[i] = m + k*math.Log(basicVariate[i]/(1-basicVariate[i]))
	}
	return logisticVariates
}
func NormalDistributionVariates(N int, m float64, sSquare float64, basicVariate []float64) []float64 {
	n := len(basicVariate) - N
	normalVariates := make([]float64, n)
	var tmp float64
	halfN := float64(N / 2)
	divider := math.Sqrt(float64(N) / 12)
	s := math.Sqrt(sSquare)
	for i := 0; i < n; i++ {
		tmp = 0.0
		for j := 0; j < N; j++ {
			tmp += basicVariate[j+i]
		}
		tmp = tmp - halfN
		tmp /= divider
		normalVariates[i] = m + tmp*s
	}
	return normalVariates
}
func SquareHiDistributionVariates(m int, variance float64, ev float64, normalDistributionVariate []float64) []float64 {
	n := len(normalDistributionVariate) - m
	squareHiVariates := make([]float64, n)
	var tmp float64
	for i := 0; i < n; i++ {
		tmp = 0.0
		for j := 0; j < m; j++ {
			tmp += math.Pow(normalDistributionVariate[j+i]-ev, 2.0) / variance
		}
		squareHiVariates[i] = tmp
	}
	return squareHiVariates
}
func FischerDistributionVariates(l float64, m float64, squareHiVariates1 []float64, squareHiVariates2 []float64) []float64 {
	n := len(squareHiVariates1)
	fischerVariates := make([]float64, n)
	for i := 0; i < n; i++ {
		fischerVariates[i] = (squareHiVariates1[i] / l) / (squareHiVariates2[i] / m)
	}
	return fischerVariates
}
