package trueskill

// Player is a struct that reflects one contestant in a game. The skill level
// is assumed to be bell shaped (normal or Gaussian distributed) with a peak
// point µ (mu) and a spread σ (sigma). The actuall skill of the player is
// assumed to be within µ +/- 3σ with 99.7% accuracy.
type Player struct {
	mu  float64
	sig float64
}

func NewPlayer(mu float64, sig float64) (p Player) {
	return Player{mu, sig}
}

func (p *Player) GetMu() float64 {
	return p.mu
}

func (p *Player) GetSigma() float64 {
	return p.sig
}

// GetCurrentSkill returns the conservative skill estimate of the player. This
// estimate is µ - 3σ. With 99.7% accuracy the actuall skill is higher than
// this value.
func (p *Player) GetCurrentSkill() float64 {
	return p.mu - 3*p.sig
}
