package trueskill

import "sort"

type playersByRank struct {
	data []Player
	rank []int
}

func (t playersByRank) Len() int { return len(t.data) }
func (t playersByRank) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
	t.rank[i], t.rank[j] = t.rank[j], t.rank[i]
}
func (t playersByRank) Less(i, j int) bool { return t.rank[i] < t.rank[j] }

// SortPlayersByRank sorts players by their rankings. The returned values
// have the same length as the provided arguments but are sorted from
// lowest (best) to highest (worst) rank.
func SortPlayersByRank(data []Player, rank []int) ([]Player, []int) {
	sort.Sort(playersByRank{data, rank})
	return data, rank
}

type teamsByRank struct {
	data []Team
	rank []int
}

func (t teamsByRank) Len() int { return len(t.data) }
func (t teamsByRank) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
	t.rank[i], t.rank[j] = t.rank[j], t.rank[i]
}
func (t teamsByRank) Less(i, j int) bool { return t.rank[i] < t.rank[j] }

// SortTeamsByRank sorts teams by their rankings. The returned values
// have the same length as the provided arguments but are sorted from
// lowest (best) to highest (worst) rank.
func SortTeamsByRank(data []Team, rank []int) ([]Team, []int) {
	sort.Sort(teamsByRank{data, rank})
	return data, rank
}
