package trueskill

import (
	"math"
	"trueskill/stats"
)

// GetDrawMargin returns the draw margin epsilon given a draw probability.
//
//  pDraw is the probability of a draw of equally skilled players [0..1]
//  beta specifies how much skill difference is defined to give a player 80%
//    chance of winning over the lower skilled player (see DefBeta).
//  nPlayers defines how many players against each other, independent of team
//    composition (e.g. 2 for chess, 3 for skat or 22 for soccer)
func GetDrawMargin(pDraw, beta float64, nPlayers uint) float64 {
	return stats.InverseCDF((pDraw+1)/2) * math.Sqrt(float64(nPlayers)) * beta
}

// GetDrawProbability returns the draw probability given a draw margin.
//
//  epsilon defines the draw margin of the game
//  beta specifies how much skill difference is defined to give a player 80%
//    chance of winning over the lower skilled player (see DefBeta).
//  nPlayers defines how many players against each other, independent of team
//    composition (e.g. 2 for chess, 3 for skat or 22 for soccer)
func GetDrawProbability(epsilon, beta float64, nPlayers uint) float64 {
	return 2 * stats.NormalCDF(epsilon/math.Sqrt(float64(nPlayers)*beta)-1)
}
