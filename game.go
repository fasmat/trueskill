package trueskill

import (
	"errors"
	"math"
)

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
//	beta defines how much difference in skill is required to expect the higher
//	  ranked player to win 80% of
//	tau specifies the increase in variance between games. This ensures that
//	  the rankings stay dynamic. Higher values will make it easier to go up
//	  and down on the ladder, lower ones make it harder.
//	pDraw specifies the probability of drawing in a random game.
func NewGame(beta, tau, pDraw float64) Game {
	return Game{
		beta:  beta,
		tau:   tau,
		pDraw: pDraw,
	}
}

// CalcNewRatings processes teams/players based on their ranks and
// updates their skills accordingly
func (g *Game) CalcNewRatings(teams []Team, ranks []int) (t []Team, err error) {
	// TODO
	_ = teams
	_ = ranks
	return
}

// CalcMatchQuality returns a value that indicates the balance of a game
// between the provided teams (0.0 = imbalanced, 1.0 = perfectly balanced)
func (g *Game) CalcMatchQuality(teams []Team) (float64, error) {
	if len(teams) > 2 {
		return 0, errors.New("does not support more than 2 teams yet")
	}

	nPlayers := float64(teams[0].Size() + teams[1].Size())

	t1Mean := teams[0].GetMu()
	t1Var := teams[0].GetVar()
	t2Mean := teams[1].GetMu()
	t2Var := teams[1].GetVar()

	sqrt := math.Sqrt(
		(nPlayers * g.beta * g.beta) /
			(nPlayers*g.beta*g.beta + t1Var + t2Var))

	exp := math.Exp(
		-1 * (t1Mean - t2Mean) * (t1Mean - t2Mean) /
			(2 * (nPlayers*g.beta*g.beta + t1Var + t2Var)))

	return sqrt * exp, nil
}
