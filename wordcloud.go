package main

import (
	"fmt"
	"github.com/sweb/wordcloud/words"
)

// Defines lower bound of word count to be displayed
const wordThreshold = 2

// Defines maximum n of n-grams
const maxNGramSize = 4

func main() {
	fmt.Println("###############################################################")
	// Contains common words
	frequencyList, err := words.PutFileContentInSlice("data/top10000de.txt")
	if err != nil {
		fmt.Println("Unable to load frequency list: " + err.Error())
	}
	wordMap := words.CountOfWords(importedWords, frequencyList)
	importedWords, err := words.PutFileContentInSlice("data/words.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Number of words: %d\n", +len(importedWords))
	fmt.Printf("Number of filtered distinct words: %d\n", +len(wordMap))
	fmt.Println("###############################################################")

	ngramMap := make(words.NGramMap)

	err = ngramMap.CombineDistinctiveNGrams(importedWords, maxNGramSize,
		wordMap, wordThreshold)
	if err != nil {
		fmt.Println("Error: n-gram evaluation skipped due to:\n" + err.Error())
	} else {
		fmt.Println("n-grams:")
		for ngram, _ := range ngramMap {
			fmt.Println(ngram)
		}
	}
	fmt.Println("###############################################################")
	sortedWordList := words.SortMapByValue(wordMap)

	for _, wordStat := range sortedWordList {
		if wordStat.Value >= wordThreshold {
			fmt.Printf("%v - %d\n", wordStat.Key, wordStat.Value)
		}
	}
}
