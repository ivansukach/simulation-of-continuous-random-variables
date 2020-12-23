package main

import (
	"github.com/ivansukach/simulation-of-continuous-random-variables/distributions"
	"github.com/ivansukach/simulation-of-continuous-random-variables/generators"
	"github.com/sirupsen/logrus"
	"math"
)

func BiasOfAnEstimator(variates []float64) (float64, float64) {
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

func main() {
	a01 := 296454621
	a02 := 302711857
	c1 := 48840859
	c2 := 37330745
	M := int(math.Pow(2, 31))
	K := 64
	n := 10000
	mNormal := 0
	N := 48
	sSquareNormal := 1
	logrus.Info("M: ", M)
	//aSequence2 := *generators.LinearCongruential(a02, c2, M, n)
	//aSequence1SpecialSize := *generators.LinearCongruential(a01, c1, M, n+K)
	//sequenceByMacLarenMarsaglia := *generators.MacLarenMarsaglia(aSequence1SpecialSize, aSequence2, K, n)
	aSequence22 := *generators.LinearCongruential(a02, c2, M, n+N)
	aSequence1SpecialSize2 := *generators.LinearCongruential(a01, c1, M, n+K+N)
	sequenceByMacLarenMarsaglia2 := *generators.MacLarenMarsaglia(aSequence1SpecialSize2, aSequence22, K, n+N)
	normalDistributionVariates := distributions.NormalDistributionVariates(N, float64(mNormal), float64(sSquareNormal), sequenceByMacLarenMarsaglia2)
	logrus.Info("First 10 variates of normal distribution")
	for i := 0; i < 10; i++ {
		logrus.Info(normalDistributionVariates[i])

	}
	logrus.Info("Expected value: ", mNormal)
	logrus.Info("Variance: ", sSquareNormal)
	unbiasedNormalEV, unbiasedNormalV := BiasOfAnEstimator(normalDistributionVariates)
	logrus.Info("Bias of an estimator expected value: ", unbiasedNormalEV)
	logrus.Info("Bias of an estimator variance: ", unbiasedNormalV)

	//x := [16]float64{1.406,0.799, 0.288, 1.010, 0.208, 1.406, 0.872, 0.671, 0.457, 0.327, 0.248, 0.327, 2.112, 1.351, 0.744, 0.669}
	//tmp_sum :=0.0
	//for i:=0; i<16; i++{
	//	tmp_sum += x[i]*math.Sqrt(x[i]*x[i]+math.Sqrt(x[i]))
	//}
	//logrus.Info("Result: ", tmp_sum/16.0)
	//binomialDistributionVariates := *distributions.BinomialDistributionVariates(mBinomial, pBinomial, sequenceByMacLarenMarsaglia2)
	//logrus.Info("First 10 variates of binomial distribution")
	//for i := 0; i < 10; i++ {
	//	logrus.Info(binomialDistributionVariates[i])
	//}
	//logrus.Info("Expected value: ", float64(mBinomial)*pBinomial)
	//logrus.Info("Variance: ", float64(mBinomial)*pBinomial*(1-pBinomial))
	//unbiasedBinomialEV, unbiasedBinomialV := BiasOfAnEstimator(binomialDistributionVariates)
	//logrus.Info("Bias of an estimator expected value: ", unbiasedBinomialEV)
	//logrus.Info("Bias of an estimator variance: ", unbiasedBinomialV)
}
