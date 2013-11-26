package trueskill

// Competitor is a participant in a game, it is either a player or a team.
type Competitor interface {
	GetCurrentSkill() float64

	GetMu() float64
	GetSigma() float64
	
	GetWinChanceAgainst(c Competitor) float64
}
