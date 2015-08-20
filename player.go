package trueskill

import "strconv"

// Player is a struct that reflects one contestant in a game. The skill level
// is assumed to be bell shaped (normal or Gaussian distributed) with a peak
// point µ (mu) and a spread σ (sigma). The actual skill of the player is
// therefore assumed to be within µ +/- 3σ with 99.7% accuracy.
type Player struct {
	id uint
	g  Gaussian
}

// NewDefaultPlayer creates and returns a Player with default values for skill
// estimation.
func NewDefaultPlayer(id uint) Player {
	return Player{
		id: id,
		g:  NewGaussian(DefMu, DefSig),
	}
}

// NewPlayer creates and returns a Player with a specified skill level.
func NewPlayer(id uint, mu, sigma float64) Player {
	return Player{
		id: id,
		g:  NewGaussian(mu, sigma),
	}
}

// GetGaussian returns the Gaussian associated with the players skill.
func (p *Player) GetGaussian() Gaussian {
	return p.g
}

// GetID returns the ID of the player
func (p *Player) GetID() uint {
	return p.id
}

// UpdateSkill updates the skill rating of player to the provided Gaussian.
func (p *Player) UpdateSkill(g Gaussian) {
	p.g = g
	return
}

// GetMu is a convenience wrapper for Gaussian.GetMu()
func (p *Player) GetMu() float64 {
	return p.g.GetMu()
}

// GetSigma is a convenience wrapper for Gaussian.GetSigma()
func (p *Player) GetSigma() float64 {
	return p.g.GetSigma()
}

func (p *Player) String() (s string) {
	g := p.GetGaussian()

	s = "Player [" + string(p.id)
	s += "] Skill-Estimate:"
	s += strconv.FormatFloat(g.GetConservativeEstimate(), 'f', 3, 64)
	s += " (mu: " + strconv.FormatFloat(g.GetMu(), 'f', 3, 64)
	s += " sig: " + strconv.FormatFloat(g.GetSigma(), 'f', 3, 64)
	s += ")"
	return
}
