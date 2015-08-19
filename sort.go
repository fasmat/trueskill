package trueskill

import "sort"

// ByRank implements sort.Interface for []Team based on rankings
type ByRank struct {
	data []Team
	rank []int
}

func (t ByRank) Len() int { return len(t.data) }
func (t ByRank) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
	t.rank[i], t.rank[j] = t.rank[j], t.rank[i]
}
func (t ByRank) Less(i, j int) bool { return t.rank[i] < t.rank[j] }

func sortTeamsByRank(data []Team, rank []int) ([]Team, []int) {
	sort.Sort(ByRank{data, rank})
	return data, rank
}
