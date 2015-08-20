package trueskill

import (
	"math"
	"testing"
)

const (
	errTolerance float64 = 1e-6
)

func TestTwoPlayerNotDraw(t *testing.T) {
	player1 := NewDefaultPlayer(1)
	player2 := NewDefaultPlayer(2)

	t.Logf("Player1: %.6f (%.6f)", player1.GetMu(), player1.GetSigma())
	t.Logf("Player2: %.6f (%.6f)", player2.GetMu(), player2.GetSigma())

	game := NewDefaultGame()

	team1 := NewTeam([]Player{player1})
	team2 := NewTeam([]Player{player2})

	quality, _ := game.CalcMatchQuality([]Team{team1, team2})
	game.CalcNewRatings([]Team{team1, team2}, []int{1, 2})

	p1mu := player1.GetMu()
	p1sig := player1.GetSigma()

	p2mu := player2.GetMu()
	p2sig := player2.GetSigma()

	if math.Abs(quality-0.447214) > errTolerance {
		t.Errorf("Expected quality to be 0.447 got %.6f", quality)
	}

	if math.Abs(p1mu-29.39583201999924) > errTolerance {
		t.Errorf("Expected p1 mu to be 29.395832 got %.6f", p1mu)
	}

	if math.Abs(p1sig-7.171475587326186) > errTolerance {
		t.Errorf("Expected p1 sigma to be 7.171476 got %.6f", p1sig)
	}

	if math.Abs(p2mu-20.60416798000076) > errTolerance {
		t.Errorf("Expected p2 mu to be 20.604168 got %.6f", p2mu)
	}

	if math.Abs(p2sig-7.171475587326186) > errTolerance {
		t.Errorf("Expected p2 sigma to be 7.171476 got %.6f", p2sig)
	}
}
