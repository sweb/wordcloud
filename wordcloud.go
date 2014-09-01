package main

import (
	"fmt"
	"github.com/sweb/wordcloud/words"
)

// Defines lower bound of word count to be displayed
var wordThreshold = 1

func main() {
	fmt.Println("###############################################################")
	importedWords, err := words.PutFileContentInSlice("data/words.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(importedWords)
	topWords, err := words.PutFileContentInSlice("data/top10000de.txt")
	if err != nil {
		fmt.Println("Unable to load frequency list: " + err.Error())
	}
	fmt.Println("###############################################################")
	wordMap := words.CountOfWords(importedWords, topWords)

	twoGrams, err := words.CreateNGrams(importedWords, 2)
	if err != nil {
		fmt.Println(err.Error())
	}

	ngramMap := make(map[string]int)

	for _, gram := range twoGrams {
		if gram.IsDistinctive(wordMap) {
			ngramMap[gram.Key()]++
		}
	}

	fmt.Println("2-Grams:")
	fmt.Println(ngramMap)
	sortedWordList := words.SortMapByValue(wordMap)

	fmt.Printf("Number of words: %d\n", +len(importedWords))
	fmt.Printf("Number of filtered distinct words: %d\n", +len(sortedWordList))
	for _, wordStat := range sortedWordList {
		if wordStat.Value >= wordThreshold {
			fmt.Printf("%v - %d\n", wordStat.Key, wordStat.Value)
		}

	}
}
