package stats

import (
	"fmt"
	"math"
)

// Gaussian represents a gaussian distribution
type Gaussian struct {
	tau float64
	pi  float64
}

// NewGaussian creates a gaussian distribution from a mean value (mu) and a
// standard deviation (sigma)
func NewGaussian(mu, sigma float64) Gaussian {
	g := Gaussian{
		pi:  1.0 / (sigma * sigma),
		tau: mu / (sigma * sigma),
	}
	return g
}

// NewGaussianFromPrecisionMean creates a gaussian distribution from a
// precision mean (tau) and a precision (pi)
func NewGaussianFromPrecisionMean(tau, pi float64) Gaussian {
	g := Gaussian{
		tau: tau,
		pi:  pi,
	}
	return g
}

// GetMu returns the mean value (mu) for a gaussian distribution
func (g *Gaussian) GetMu() float64 {
	return g.tau / g.pi
}

// GetSigma returns the standard deviation of a gaussian distribution
func (g *Gaussian) GetSigma() float64 {
	return math.Sqrt(1.0 / g.pi)
}

// GetVar returns the variance of a gaussian distribution
func (g *Gaussian) GetVar() float64 {
	return g.GetSigma() * g.GetSigma()
}

// GetNormalizationConstant returns the normalization constant for the gaussian
// distribution.
func (g *Gaussian) GetNormalizationConstant() float64 {
	return 1.0 / (math.Sqrt(2*math.Pi) * g.GetSigma())
}

// GetConservativeEstimate returns the conservative skill estimate of the
// player. This estimate is µ - 3s. With 99.7% accuracy the actual skill is
// higher than this value.
func (g *Gaussian) GetConservativeEstimate() float64 {
	return g.GetMu() - 3*g.GetSigma()
}

// String can be used to print a Gaussian
func (g *Gaussian) String() string {
	text := fmt.Sprintf("μ=%2.4f, σ=%2.4f", g.GetMu(), g.GetSigma())
	return text
}

// MultiplyGaussian multiplies two Gaussian distributions and returns the
// result as a new Gaussian
func MultiplyGaussian(fac1, fac2 Gaussian) Gaussian {
	prod := Gaussian{
		tau: fac1.tau + fac2.tau,
		pi:  fac1.pi + fac2.pi,
	}

	return prod
}

// DivideGaussian divides two Gaussian distributions and returns the result as
// a new Gaussian
func DivideGaussian(numerator, denominator Gaussian) Gaussian {
	quot := Gaussian{
		tau: numerator.tau - denominator.tau,
		pi:  numerator.pi - denominator.pi,
	}
	return quot
}

// LogProductNormalisation computes the log-normalisation factor when two
// normalised Gaussians get multiplied
func LogProductNormalisation(left, right Gaussian) float64 {
	if left.pi == 0.0 || right.pi == 0.0 {
		return 0.0
	}
	varSum := left.GetVar() + right.GetVar()
	muDiff := left.GetMu() - right.GetMu()
	logSqrt2Pi := math.Log(math.Sqrt(2 * math.Pi))

	return -logSqrt2Pi - math.Log(varSum/2.0) - (muDiff*muDiff)/(2.0*varSum)
}

// LogRatioNormalisation computes the log-normalisation factor when two
// normalised Gaussians get divided
func LogRatioNormalisation(numerator, denominator Gaussian) float64 {
	if numerator.pi == 0.0 || denominator.pi == 0.0 {
		return 0.0
	}

	varDiff := denominator.GetVar() - numerator.GetVar()
	muDiff := numerator.GetMu() - denominator.GetMu()

	logSqrt2Pi := math.Log(math.Sqrt(2 * math.Pi))
	return math.Log(denominator.GetVar()) + logSqrt2Pi -
		math.Log(varDiff)/2.0 + (muDiff*muDiff)/(2.0*varDiff)
}

// AbsoluteDiff calculates the absolute difference between two gaussian
// distributions
func AbsoluteDiff(left, right Gaussian) float64 {
	result := math.Max(math.Abs(left.tau-right.tau),
		math.Sqrt(math.Abs(left.pi-right.pi)))

	return result
}
