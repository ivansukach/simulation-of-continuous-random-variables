package distributions

import "math"

func one(x float64) float64 {
	if x <= 0 {
		return 0.0
	}
	return 1.0
}
func BinomialDistributionVariates(m int, p float64, basicVariate []float64) *[]float64 {
	n := len(basicVariate) - m
	binomialVariates := make([]float64, n)
	var tmp float64
	for i := 0; i < n; i++ {
		tmp = 0.0
		for j := 0; j < m; j++ {
			tmp += one(p - basicVariate[j+i])
		}
		binomialVariates[i] = tmp
	}
	return &binomialVariates
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
