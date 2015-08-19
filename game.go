package trueskill

const (
	// DefMu specifies the default mean value (mu) of a new player
	DefMu = 25.0

	// DefSig specifies the default standard deviation (sigma) of a new
	// player
	DefSig = DefMu / 3

	// DefBeta defines the default beta for games. This specifies how much
	// difference in skill relates to a 80% chance to win for the higher
	// ranked player.
	DefBeta = DefSig / 2

	// DefTau specifies the increase in variance between games. The higher
	// this value the easier it is to go up and down on the ladder.
	// 1% of the default sigma gives reasonable dynamics.
	DefTau = DefSig / 100
)

// Game holds information about a game where rankings shall be calculated for
type Game struct {
	beta  float64
	tau   float64
	pDraw float64
}

// NewDefaultGame creates a new game with default properties
func NewDefaultGame() Game {
	return NewGame(DefBeta, DefTau, 0)
}

// NewGame creates a new game with the specified properties
//
//  beta defines how much difference in skill is required to expect the higher
//    ranked player to win 80% of
//  tau specifies the increase in variance between games. This ensures that
//    the rankings stay dynamic. Higher values will make it easier to go up
//    and down on the ladder, lower ones make it harder.
//  pDraw specifies the probability of drawing in a random game.
func NewGame(beta, tau, pDraw float64) Game {
	return Game{
		beta:  beta,
		tau:   tau,
		pDraw: pDraw,
	}
}
}
