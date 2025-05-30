package trueskill

import "testing"

func TestSortPlayersByRank(t *testing.T) {
	t.Parallel()

	players := generatePlayers(5)

	dataIn := []Player{
		players[0],
		players[4],
		players[3],
		players[1],
		players[2],
	}

	rankIn := []int{0, 4, 3, 1, 2}

	data, rank := SortPlayersByRank(dataIn, rankIn)

	if len(data) != len(players) {
		t.Error("Lost data values during sorting. Expected",
			len(players), "got", len(data))
	}

	if len(rank) != len(rankIn) {
		t.Error("Lost rank values during sorting. Expected",
			len(rankIn), "got", len(rank))
	}

	for i, v := range data {
		if v.GetID() != i {
			t.Error("Element", i, "of players is not the same as expected")
		}

		if rank[i] != i {
			t.Error("Element", i, "of rank is not the same as expected")
		}
	}
}

func TestSortTeamsByRank(t *testing.T) {
	t.Parallel()

	players := generatePlayers(5)
	dataIn := []Team{
		NewTeam([]Player{players[4]}),
		NewTeam([]Player{players[3]}),
		NewTeam([]Player{players[2]}),
		NewTeam([]Player{players[1]}),
		NewTeam([]Player{players[0]}),
	}
	rankIn := []int{5, 4, 3, 1, 1}

	dataOut := []Team{
		NewTeam([]Player{players[1]}),
		NewTeam([]Player{players[0]}),
		NewTeam([]Player{players[2]}),
		NewTeam([]Player{players[3]}),
		NewTeam([]Player{players[4]}),
	}
	rankOut := []int{1, 1, 3, 4, 5}

	data, rank := SortTeamsByRank(dataIn, rankIn)

	if len(data) != len(dataOut) {
		t.Error("Lost data values during sorting. Expected",
			len(dataOut), "got", len(data))
	}

	if len(rank) != len(rankOut) {
		t.Error("Lost rank values during sorting. Expected",
			len(rankOut), "got", len(rank))
	}

	for i, v := range data {
		if v.String() != dataOut[i].String() {
			t.Error("Element", i, "of teams is not the same as expected")
		}
	}

	for i, v := range rank {
		if v != rankOut[i] {
			t.Error("Element", i, "of rank is not the same as expected")
		}
	}
}

func generatePlayers(n int) []Player {
	p := make([]Player, n)

	for i := 0; i < n; i++ {
		p[i] = NewDefaultPlayer(i)
	}
	return p
}
