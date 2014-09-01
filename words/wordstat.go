package words

import (
	"sort"
	"strings"
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

// A data structure to hold n consecutive words
type NGram []string

// Concatinates all elements of the n-gram in order to create a key representing
// it.
func (ng NGram) Key() string {
	// n-grams always have at least 1 element, thus this should always work
	key := ng[0]
	for i := 1; i < len(ng); i++ {
		key += " " + ng[i]
	}
	return key
}

// Determines if the n-gram is distinctive. In this case this means that it only
// contains relevant words. Only relevant words are left after CountOfWords()
// was used to create the word list of distinct words.
func (ng NGram) IsDistinctive(distinctWords map[string]int) bool {
	distinctive := true
	for _, word := range ng {
		if distinctWords[strings.ToLower(word)] == 0 {
			return false
		}
	}
	return distinctive
}
