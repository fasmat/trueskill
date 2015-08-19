package trueskill

import "testing"

func TestSortByRank(t *testing.T) {
	dataIn := []Team{NewEmptyTeam(), NewEmptyTeam(), NewEmptyTeam(),
		NewEmptyTeam()}
	rankIn := []int{4, 3, 1, 1}

	dataOut := []Team{NewEmptyTeam(), NewEmptyTeam(), NewEmptyTeam(),
		NewEmptyTeam()}
	rankOut := []int{1, 1, 3, 4}

	data, rank := sortTeamByRank(dataIn, rankIn)

	if len(data) != len(dataOut) {
		t.Error("Lost data values during sorting. Expected",
			len(dataOut), "got", len(data))
	}

	if len(rank) != len(rankOut) {
		t.Error("Lost rank values during sorting. Expected",
			len(rankOut), "got", len(rank))
	}

	for i, v := range rank {
		if v != rankOut[i] {
			t.Error("Element", i, "of rank is not the same as expected")
		}
	}
}
