package trueskill

// Team is a composition of players that play together. The skill of a team
// (µ and σ) is determined by the skills of the players that form the team.
type Team struct {
	players []Player
}

// NewEmptyTeam creates a team without any players. To add players use
// AddPlayers.
func NewEmptyTeam() Team {
	return Team{
		players: make([]Player, 0),
	}
}

// NewTeam creates a team from a slice of players.
func NewTeam(p []Player) (t Team) {
	t.players = p
	return
}

// Size returns the number of players in the team
func (t *Team) Size() int {
	return len(t.players)
}

// AddPlayer adds a single player to the team.
func (t *Team) AddPlayer(p Player) {
	t.players = append(t.players, p)
	return
}

// AddPlayers adds players to the team.
func (t *Team) AddPlayers(p []Player) {
	t.players = append(t.players, p...)
	return
}

// GetPlayers returns the players the team is composed of.
func (t *Team) GetPlayers() (p []Player) {
	return t.players
}

// TODO: Add functions to remove players from a team

// GetMu returns the sum of all means of the team
func (t *Team) GetMu() (sum float64) {
	for _, p := range t.players {
		sum += p.GetMu()
	}
	return
}

// GetVar returns the combined variance of the team
func (t *Team) GetVar() (sum float64) {
	for _, p := range t.players {
		sum += p.GetSigma() * p.GetSigma()
	}
	return
}
