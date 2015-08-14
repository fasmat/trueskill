package trueskill

import "math"

// Team is a composition of players that play together. The skill of a team
// (µ and σ) is determined by the skills of the players that form the team.
type Team struct {
	players []Player
}

// NewTeam creates a team from a slice of players.
func NewTeam(p []Player) (t Team) {
	t.players = p
	t.updateSkill()
	return
}

// GetCurrentSkill returns the conservative skill estimate of the team. This
// estimate is µ - 3σ. With 99.7% accuracy the actuall skill is higher than
// this value.
func (t *Team) GetCurrentSkill() float64 {
	return t.mu - 3*t.sig
}

// updateSkill must be called everytime player composition changes or skills of
// players change (e.g. after a game).
func (t *Team) updateSkill() {
	for _, p := range t.players {
		t.mu += p.mu
		t.sig += p.sig * p.sig
	}

	t.sig = math.Sqrt(t.sig)
	return
}
