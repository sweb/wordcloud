package words

import (
	"errors"
	"strings"
)

// A data structure to hold n consecutive words
type NGram []string

type NGramMap map[string]int

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
func (ng NGram) IsDistinctive(distinctWords map[string]int, threshold int) bool {
	distinctive := true
	for _, word := range ng {
		if distinctWords[strings.ToLower(word)] < threshold {
			return false
		}
	}
	return distinctive
}

// This function creates n-grams based on a given set of words and the size of
// n. n needs to be greater than 0.
func createNGrams(words []string, n int) ([]NGram, error) {
	if n < 0 {
		return nil, errors.New("n of n-gram needs to be > 0")
	}
	// n reduces total number of n-grams but even if n equals the number of
	// words, one n-gram is created.
	numberOfNGrams := len(words) - (n - 1)
	nGramList := make([]NGram, numberOfNGrams)
	for i := 0; i < numberOfNGrams; i++ {
		nGram := make(NGram, n)
		for j := 0; j < n; j++ {
			nGram[j] = words[i+j]
		}
		nGramList[i] = nGram
	}
	return nGramList, nil
}

// Adds a list of words in chosen n-gram representation to a map containing all
// n-grams that are distinctive according to a map of distinctive words.
func (nGramMap NGramMap) addDistinctiveNGrams(words []string, n int,
	distinctiveWords map[string]int, threshold int) error {
	nGrams, err := createNGrams(words, n)
	if err != nil {
		return err
	}
	for _, gram := range nGrams {
		if gram.IsDistinctive(distinctiveWords, threshold) {
			nGramMap[gram.Key()]++
		}
	}
	return nil
}

// Adds n-grams to a map of distinctive n-grams and counts the number of
// occurences in a given set of words. n is increased starting from 2 to a given
// maximum value.
func (nGramMap NGramMap) CombineDistinctiveNGrams(words []string, maxN int,
	distinctiveWords map[string]int, threshold int) error {
	for i := 2; i <= maxN; i++ {
		err := nGramMap.addDistinctiveNGrams(words, i, distinctiveWords, threshold)
		if err != nil {
			return err
		}
	}
	return nil
}
