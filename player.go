package trueskill

const (
	// DefMu specifies the default mean value (mu) of a new player
	DefMu float64 = 25

	// DefSig specifies the default standard deviation (sigma) of a new
	// player
	DefSig float64 = 25 / 3
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

// GetCurrentSkill returns the conservative skill estimate of the player. This
// estimate is µ - 3σ. With 99.7% accuracy the actual skill is higher than
// this value.
func (p *Player) GetCurrentSkill() float64 {
	return p.g.GetMu() - 3*p.g.GetSigma()
}
