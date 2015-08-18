package trueskill

import "math"

var (
	// DefMu specifies the default mean value (mu) of a new player
	DefMu = 25.0

	// DefSig specifies the default standard deviation (sigma) of a new
	// player
	DefSig = math.Sqrt(DefMu)
)

// Player is a struct that reflects one contestant in a game. The skill level
// is assumed to be bell shaped (normal or Gaussian distributed) with a peak
// point µ (mu) and a spread σ (sigma). The actual skill of the player is
// therefore assumed to be within µ +/- 3σ with 99.7% accuracy.
type Player struct {
	g Gaussian
}

// NewDefaultPlayer creates and returns a Player with default values for skill
// estimation.
func NewDefaultPlayer() (p Player) {
	return Player{NewGaussian(DefMu, DefSig)}
}

// NewPlayer creates and returns a Player with a specified skill level.
func NewPlayer(mu, sigma float64) (p Player) {
	p.g = NewGaussian(mu, sigma)
	return
}
