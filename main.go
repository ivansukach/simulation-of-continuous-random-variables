package main

import (
	"github.com/ivansukach/simulation-of-continuous-random-variables/distributions"
	"github.com/ivansukach/simulation-of-continuous-random-variables/generators"
	"github.com/sirupsen/logrus"
	"math"
)

func EmpiricalEstimates(variates []float64) (float64, float64) {
	n := len(variates)
	logrus.Info("n=", n)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += variates[i]
	}
	E := sum / float64(n)
	sum = 0.0
	for i := 0; i < n; i++ {
		sum += math.Pow(variates[i]-E, 2)
	}
	D := sum / float64(n-1)
	return E, D
}
func FischerVariance(l float64, m float64) float64 {
	if m > 4 {
		return 2 * m * m * (l + m - 2) / (l * (m - 2) * (m - 2) * (m - 4))
	} else {
		return 0.0
	}

}
func normalDistributionInfo(variates []float64) {
	n := len(variates)
	p := [6]int{0, 0, 0, 0, 0, 0}
	for i := 0; i < n; i++ {
		if variates[i] < -2.0 && variates[i] >= -3.0 {
			p[0] += 1
		} else if variates[i] < -1.0 && variates[i] >= -2.0 {
			p[1] += 1
		} else if variates[i] < 0.0 && variates[i] >= -1.0 {
			p[2] += 1
		} else if variates[i] < 1.0 && variates[i] >= 0.0 {
			p[3] += 1
		} else if variates[i] < 2.0 && variates[i] >= 1.0 {
			p[4] += 1
		} else if variates[i] < 3.0 && variates[i] >= 2.0 {
			p[5] += 1
		}
	}
	logrus.Info("-3 <= x < -2 :", p[0])
	logrus.Info("-2 <= x < -1 :", p[1])
	logrus.Info("-1 <= x < 0 :", p[2])
	logrus.Info("0 <= x < 1 :", p[3])
	logrus.Info("1 <= x < 2 :", p[4])
	logrus.Info("2 <= x < 3 :", p[5])
}
func main() {
	a01 := 296454621
	a02 := 302711857
	c1 := 48840859
	c2 := 37330745
	M := int(math.Pow(2, 31))
	K := 64
	n := 10000
	mNormal := 0
	mLogistic := 2.0
	kLogistic := 3.0
	lFischer := 5
	mFischer := 3
	N := 48
	sSquareNormal := 1
	logrus.Info("M: ", M)
	aSequence2 := *generators.LinearCongruential(a02, c2, M, n)
	aSequence1SpecialSize := *generators.LinearCongruential(a01, c1, M, n+K)
	sequenceByMacLarenMarsaglia := *generators.MacLarenMarsaglia(aSequence1SpecialSize, aSequence2, K, n)
	aSequence22 := *generators.LinearCongruential(a02, c2, M, n+N)
	aSequence1SpecialSize2 := *generators.LinearCongruential(a01, c1, M, n+K+N)
	sequenceByMacLarenMarsaglia2 := *generators.MacLarenMarsaglia(aSequence1SpecialSize2, aSequence22, K, n+N)
	aSequence23 := *generators.LinearCongruential(a02, c2, M, n+lFischer+N)
	aSequence1SpecialSize3 := *generators.LinearCongruential(a01, c1, M, n+K+lFischer+N)
	sequenceByMacLarenMarsaglia3 := *generators.MacLarenMarsaglia(aSequence1SpecialSize3, aSequence23, K, n+lFischer+N)
	aSequence24 := *generators.LinearCongruential(a02, c2, M, n+mFischer+N)
	aSequence1SpecialSize4 := *generators.LinearCongruential(a01, c1, M, n+K+mFischer+N)
	sequenceByMacLarenMarsaglia4 := *generators.MacLarenMarsaglia(aSequence1SpecialSize4, aSequence24, K, n+mFischer+N)

	normalDistributionVariates := distributions.NormalDistributionVariates(N, float64(mNormal), float64(sSquareNormal), sequenceByMacLarenMarsaglia2)
	normalDistributionVariates2 := distributions.NormalDistributionVariates(N, float64(mNormal), float64(sSquareNormal), sequenceByMacLarenMarsaglia3)
	normalDistributionVariates3 := distributions.NormalDistributionVariates(N, float64(mNormal), float64(sSquareNormal), sequenceByMacLarenMarsaglia4)
	logrus.Info("First 10 variates of normal distribution")
	for i := 0; i < 10; i++ {
		logrus.Info(normalDistributionVariates[i])

	}

	logrus.Info("Expected value: ", mNormal)
	logrus.Info("Variance: ", sSquareNormal)
	empiricalNormalEV, empiricalNormalV := EmpiricalEstimates(normalDistributionVariates)
	logrus.Info("Empirical estimates of expected value: ", empiricalNormalEV)
	logrus.Info("Empirical estimates of variance: ", empiricalNormalV)

	logisticDistributionVariates := distributions.LogisticDistributionVariates(mLogistic, kLogistic, sequenceByMacLarenMarsaglia)
	logrus.Info("First 10 variates of logistic distribution")
	for i := 0; i < 10; i++ {
		logrus.Info(logisticDistributionVariates[i])

	}
	logrus.Info("Expected value: ", mLogistic)
	logrus.Info("Variance: ", math.Pow(kLogistic/math.Sqrt(3)*math.Pi, 2.0))
	empiricalLogisticEV, empiricalLogisticV := EmpiricalEstimates(logisticDistributionVariates)
	logrus.Info("Empirical estimates of expected value: ", empiricalLogisticEV)
	logrus.Info("Empirical estimates of variance: ", empiricalLogisticV)

	logrus.Info("Normal distribution 1 for square hi")
	normalDistributionInfo(normalDistributionVariates2)
	logrus.Info("Normal distribution 2 for square hi")
	normalDistributionInfo(normalDistributionVariates3)
	empiricalNormal2EV, empiricalNormal2V := EmpiricalEstimates(normalDistributionVariates2)
	logrus.Info("Empirical estimates of normal expected value: ", empiricalNormal2EV)
	logrus.Info("Empirical estimates of normal variance: ", empiricalNormal2V)
	empiricalNormal3EV, empiricalNormal3V := EmpiricalEstimates(normalDistributionVariates3)
	logrus.Info("Empirical estimates of normal expected value: ", empiricalNormal3EV)
	logrus.Info("Empirical estimates of normal variance: ", empiricalNormal3V)

	squareHiVariates1 := distributions.SquareHiDistributionVariates(lFischer, empiricalNormal2V, empiricalNormal2EV, normalDistributionVariates2)
	squareHiVariates2 := distributions.SquareHiDistributionVariates(mFischer, empiricalNormal3V, empiricalNormal3EV, normalDistributionVariates3)

	//warning: big variate of x^2 distribution
	empiricalSquareHi1EV, empiricalSquareHi1V := EmpiricalEstimates(squareHiVariates1)
	logrus.Info("Empirical estimates of square-hi expected value: ", empiricalSquareHi1EV)
	logrus.Info("Empirical estimates of square-hi variance: ", empiricalSquareHi1V)

	fischerDistributionVariates := distributions.FischerDistributionVariates(float64(lFischer), float64(mFischer), squareHiVariates1, squareHiVariates2)
	logrus.Info("First 10 variates of fischer distribution")
	for i := 0; i < 10; i++ {
		logrus.Info(fischerDistributionVariates[i])
	}
	logrus.Info("Expected value: ", float64(mFischer)/float64(mFischer-2))
	logrus.Info("M should be more then 4")
	logrus.Info("Variance: ", FischerVariance(float64(lFischer), float64(mFischer)))
	logrus.Info("Variance reverse(l, m): ", FischerVariance(float64(mFischer), float64(lFischer)))
	empiricalFischerEV, empiricalFischerV := EmpiricalEstimates(fischerDistributionVariates)
	logrus.Info("Empirical estimates of expected value: ", empiricalFischerEV)
	logrus.Info("Empirical estimates of variance: ", empiricalFischerV)
}
