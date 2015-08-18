package trueskill

import "math"

// Gaussian represents a gaussian distribution
type Gaussian struct {
	tau float64
	pi  float64
}

// NewGaussian creates a gaussian distribution from a mean value (mu) and a
// standard deviation (sigma)
func NewGaussian(mu, sigma float64) Gaussian {
	g := Gaussian{
		pi:  1 / sigma * sigma,
		tau: mu / sigma * sigma,
	}
	return g
}

// GetMu returns the mean value (mu) for a gaussian distribution
func (g *Gaussian) GetMu() float64 {
	return g.tau / g.pi
}

// GetSigma returns the standard deviation of a gaussian distribution
func (g *Gaussian) GetSigma() float64 {
	return math.Sqrt(1 / g.pi)
}

// GetConservativeEstimate returns the conservative skill estimate of the
// player. This estimate is µ - 3σ. With 99.7% accuracy the actual skill is
// higher than this value.
func (g *Gaussian) GetConservativeEstimate() float64 {
	return g.GetMu() - 3*g.GetSigma()
}

// Mul multiplies two Gaussian distributions and returns the result as a new
// Gaussian
func (g *Gaussian) Mul(fac Gaussian) Gaussian {
	prod := Gaussian{
		tau: g.tau + fac.tau,
		pi:  g.pi + fac.pi,
	}

	return prod
}

// Div divides two Gaussian distributions and returns the result as a new
// Gaussian
func (g *Gaussian) Div(div Gaussian) Gaussian {
	quot := Gaussian{
		tau: g.tau - div.tau,
		pi:  g.pi - div.pi,
	}
	return quot
}
