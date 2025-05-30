package trueskill

import (
	"fmt"

	"github.com/fasmat/trueskill/stats"
)

// Player is a struct that reflects one contestant in a game. The skill level
// is assumed to be bell shaped (normal or Gaussian distributed) with a peak
// point µ (mu) and a spread σ (sigma). The actual skill of the player is
// therefore assumed to be within µ +/- 3σ with 99.7% accuracy.
type Player struct {
	id int
	g  stats.Gaussian
}

// NewDefaultPlayer creates and returns a Player with default values for skill
// estimation.
func NewDefaultPlayer(id int) Player {
	return Player{
		id: id,
		g:  stats.NewGaussian(DefMu, DefSig),
	}
}

// NewPlayer creates and returns a Player with a specified skill level.
func NewPlayer(id int, mu, sigma float64) Player {
	return Player{
		id: id,
		g:  stats.NewGaussian(mu, sigma),
	}
}

// GetGaussian returns the Gaussian associated with the players skill.
func (p *Player) GetGaussian() stats.Gaussian {
	return p.g
}

// GetID returns the ID of the player
func (p *Player) GetID() int {
	return p.id
}

// GetSkill returns the skill of the player as single number
func (p *Player) GetSkill() float64 {
	g := p.GetGaussian()
	return g.GetConservativeEstimate()
}

// UpdateSkill updates the skill rating of player to the provided Gaussian.
func (p *Player) UpdateSkill(g stats.Gaussian) {
	p.g = g
}

// GetMu is a convenience wrapper for Gaussian.GetMu()
func (p *Player) GetMu() float64 {
	return p.g.GetMu()
}

// GetSigma is a convenience wrapper for Gaussian.GetSigma()
func (p *Player) GetSigma() float64 {
	return p.g.GetSigma()
}

// GetVar is a convenience wrapper for Gaussian.GetVar()
func (p *Player) GetVar() float64 {
	return p.g.GetVar()
}

func (p *Player) String() string {
	return fmt.Sprintf("Player [%d] Skill-Estimate: %2.4f (μ=%2.4f, σ=%2.4f)",
		p.id,
		p.GetSkill(),
		p.GetMu(),
		p.GetSigma(),
	)
}
