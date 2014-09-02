package words

import (
	"sort"
)

// A data structure to hold a key/value pair for word statistics.
type WordStat struct {
	Key   string
	Value int
}

// A slice of WordStatistics that implements sort.Interface to sort by Value.
type WordStatList []WordStat

func (p WordStatList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p WordStatList) Len() int      { return len(p) }

// Attention: Less is used as Greater in order to have a descending list
func (p WordStatList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a WordStatList, then sort and return it.
func SortMapByValue(m map[string]int) WordStatList {
	p := make(WordStatList, len(m))
	i := 0
	for k, v := range m {
		p[i] = WordStat{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
